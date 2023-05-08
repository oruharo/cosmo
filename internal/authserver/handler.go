package authserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	cosmov1alpha1 "github.com/cosmo-workspace/cosmo/api/v1alpha1"
	"github.com/cosmo-workspace/cosmo/pkg/auth/session"
	"github.com/cosmo-workspace/cosmo/pkg/clog"
	"github.com/cosmo-workspace/cosmo/pkg/kosmo"
	authv1alpha1 "github.com/cosmo-workspace/cosmo/proto/gen/auth-server/v1alpha1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) auth(rw http.ResponseWriter, req *http.Request) {
	log := clog.FromContext(req.Context()).WithCaller()

	log.Info("------ auth start ------------")
	displayHeaders(req.Header)

	// Bypass manifest.json not to check session. By default, manifest.json is requested without cookie.
	// https://developer.mozilla.org/en-US/docs/Web/Manifest
	if strings.Contains(strings.ToLower(req.URL.Path), "/manifest.json") {
		return
	}
	ses, err := s.sessionStore.Get(req, s.SessionName)
	if ses == nil || err != nil {
		log.Error(err, "failed to get session from store. err=%s")
		rw.WriteHeader(http.StatusUnauthorized)
		s.redirectToLoginPage(rw, req)
		return
	}
	if ses.IsNew {
		log.Info("not authorized")
		rw.WriteHeader(http.StatusUnauthorized)
		s.redirectToLoginPage(rw, req)
		return
	}

	sesInfo := session.Get(ses)
	log.Debug().Info("get session.", " sesInfo=", sesInfo)

	// check user name is owner's
	userName := req.Header.Get("X-Cosmo-UserName")
	if userName != "" && sesInfo.UserName != userName {
		log.Debug().Info("invalid authorization.", " storedUserName=", sesInfo.UserName, " ownerName=", userName)
		s.redirectToLoginPage(rw, req)
		return
	}

	log.Debug().Info("authorized.", " path=", req.URL.Path)
}

func displayHeaders(headers http.Header) {
	fmt.Println("headers: {")
	for key, value := range headers {
		fmt.Println("    " + key + ": " + value[0])
	}
	fmt.Println("}")
	fmt.Println()
}

func (s Server) redirectToLoginPage(w http.ResponseWriter, r *http.Request) {

	// q := make(url.Values)
	// q.Add("redirect_to", "https://whoami3-k3d-aa-cosmo2-hashiro.cosmo4.goblab51.com/")
	// redirectURL, _ := url.Parse("https://p3000-aa-cosmo2-hashiro.cosmo4.goblab51.com/#/signin?" + q.Encode())
	// http.Redirect(w, r, redirectURL.String(), http.StatusFound)

	// redirectURL := "https://p3000-aa-cosmo2-hashiro.cosmo4.goblab51.com/#/signin"
	redirectURL := s.RedirectUrl

	bodyFormat := `
<!DOCTYPE html>
<html lang="en">
	<head>
		<title>COSMO Dashboard redirector</title>
		<script type="module">
			const originalUrl = encodeURIComponent(window.location.href);
			const redirectUrl = "%s" + "?redirect_to=" + originalUrl;
			window.location.href = redirectUrl;
			console.log(redirectUrl)
		</script>
	</head>
	<body>redirect to</body>
</html>
`
	body := fmt.Sprintf(bodyFormat, redirectURL)
	w.Write([]byte(body))
}

type ctxKeyResponseWriter struct{}
type ctxKeyRequest struct{}

func (s *Server) handlerArgsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, ctxKeyResponseWriter{}, w)
		ctx = context.WithValue(ctx, ctxKeyRequest{}, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handlerArgsFromContext(ctx context.Context) (*http.Request, http.ResponseWriter) {
	r := ctx.Value(ctxKeyRequest{}).(*http.Request)
	w := ctx.Value(ctxKeyResponseWriter{}).(http.ResponseWriter)
	return r, w
}

// ---------------------------------------------------
func (s *Server) verifyAndGetLoginUser(ctx context.Context) (loginUser *cosmov1alpha1.User, deadline time.Time, err error) {
	r, _ := handlerArgsFromContext(ctx)
	if r.Header.Get("Cookie") == "" {
		return nil, deadline, kosmo.NewUnauthorizedError("session is not found", err)
	}
	ses, err := s.sessionStore.Get(r, s.SessionName)
	if ses == nil || err != nil {
		return nil, deadline, kosmo.NewUnauthorizedError("failed to get session from store", err)
	}
	if ses.IsNew {
		return nil, deadline, kosmo.NewUnauthorizedError("session is invarild", err)
	}

	sesInfo := session.Get(ses)

	userName := sesInfo.UserName
	if userName == "" {
		return nil, deadline, kosmo.NewInternalServerError("userName is empty", nil)
	}

	deadline = time.Unix(sesInfo.Deadline, 0)
	if deadline.Before(time.Now()) {
		return nil, deadline,
			kosmo.NewUnauthorizedError(fmt.Sprintf("deadline is before the current time: deadline %v", deadline), nil)
	}

	loginUser, err = s.Klient.GetUser(ctx, userName)
	if err != nil {
		return nil, deadline, err
	}

	return loginUser, deadline, nil
}

func (s *Server) Verify(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[authv1alpha1.VerifyResponse], error) {
	// log := s.Log.WithCaller()
	log := clog.FromContext(ctx).WithCaller()

	loginUser, deadline, err := s.verifyAndGetLoginUser(ctx)
	if err != nil {
		return nil, ErrResponse(log, err)
	}

	return connect.NewResponse(&authv1alpha1.VerifyResponse{
		UserName:              loginUser.Name,
		ExpireAt:              timestamppb.New(deadline),
		RequirePasswordUpdate: false,
	}), nil
}

func (s *Server) Login(ctx context.Context, req *connect.Request[authv1alpha1.LoginRequest]) (*connect.Response[authv1alpha1.LoginResponse], error) {
	log := clog.FromContext(ctx).WithCaller()
	log.Debug().Info("request", "username", req.Msg.UserName)

	r, w := handlerArgsFromContext(ctx)

	// Check name
	user, err := s.Klient.GetUser(ctx, req.Msg.UserName)
	if err != nil {
		log.Info(err.Error(), "username", req.Msg.UserName)
		return nil, ErrResponse(log, kosmo.NewForbiddenError("incorrect user or password", nil))
	}
	// Check password
	authrizer, ok := s.Authorizers[user.Spec.AuthType]
	if !ok {
		log.Info("authrizer not found", "username", req.Msg.UserName, "authType", user.Spec.AuthType)
		return nil, ErrResponse(log, kosmo.NewServiceUnavailableError("incorrect user or password", nil))
	}
	verified, err := authrizer.Authorize(ctx, req.Msg)
	if err != nil {
		log.Error(err, "authorize failed", "username", req.Msg.UserName)
		return nil, ErrResponse(log, kosmo.NewForbiddenError("incorrect user or password", nil))

	}
	if !verified {
		log.Info("login failed: password invalid", "username", req.Msg.UserName)
		return nil, ErrResponse(log, kosmo.NewForbiddenError("incorrect user or password", nil))
	}
	var isDefault bool
	if cosmov1alpha1.UserAuthType(user.Spec.AuthType) == cosmov1alpha1.UserAuthTypePasswordSecert {
		isDefault, err = s.Klient.IsDefaultPassword(ctx, req.Msg.UserName)
		if err != nil {
			log.Error(err, "failed to check is default password", "username", req.Msg.UserName)
			return nil, ErrResponse(log, kosmo.NewInternalServerError("", nil))
		}
	}

	// Create session
	now := time.Now()
	expireAt := now.Add(time.Duration(s.MaxAgeSeconds) * time.Second)

	ses, _ := s.sessionStore.New(r, s.SessionName)
	sesInfo := session.Info{
		UserName: req.Msg.UserName,
		Deadline: expireAt.Unix(),
	}
	log.DebugAll().Info("save session", "sesInfo", sesInfo)
	ses = session.Set(ses, sesInfo)
	err = ses.Save(r, w)
	if err != nil {
		log.Error(err, "failed to save session")
		return nil, ErrResponse(log, err)
	}

	return connect.NewResponse(&authv1alpha1.LoginResponse{
		UserName:              req.Msg.UserName,
		ExpireAt:              timestamppb.New(expireAt),
		RequirePasswordUpdate: isDefault,
	}), nil
}

func (s *Server) Logout(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	log := clog.FromContext(ctx).WithCaller()

	_, _, err := s.verifyAndGetLoginUser(ctx)
	if err != nil {
		return nil, ErrResponse(log, err)
	}

	// clear session
	cookie := s.sessionOption()
	cookie.MaxAge = -1
	_, w := handlerArgsFromContext(ctx)
	http.SetCookie(w, cookie)

	resp := connect.NewResponse(&emptypb.Empty{})

	return resp, nil
}
