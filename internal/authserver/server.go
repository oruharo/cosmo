package authserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/sessions"

	cosmov1alpha1 "github.com/cosmo-workspace/cosmo/api/v1alpha1"
	"github.com/cosmo-workspace/cosmo/pkg/auth"
	"github.com/cosmo-workspace/cosmo/pkg/auth/session"
	"github.com/cosmo-workspace/cosmo/pkg/clog"
	"github.com/cosmo-workspace/cosmo/pkg/kosmo"
	"github.com/cosmo-workspace/cosmo/proto/gen/auth-server/v1alpha1/authserverv1alpha1connect"
)

// Server serves dashboard APIs and UI static files
// It implements https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/manager#Runnable
type Server struct {
	Log                 *clog.Logger
	Klient              kosmo.Client
	GracefulShutdownDur time.Duration
	ResponseTimeout     time.Duration
	StaticFileDir       string
	Port                int
	MaxAgeSeconds       int
	SessionName         string
	CookieDomain        string
	CookieHashKey       string
	CookieBlockKey      string
	RedirectUrl         string
	TLSPrivateKeyPath   string
	TLSCertPath         string
	Insecure            bool

	Authorizers  map[cosmov1alpha1.UserAuthType]auth.Authorizer
	http         *http.Server
	sessionStore sessions.Store
}

func (s *Server) setupRouter() *http.ServeMux {

	mux := http.NewServeMux()

	// forwordauth
	mux.HandleFunc("/auth", s.auth)

	// grpc
	path, authServerHandler := authserverv1alpha1connect.NewAuthServiceHandler(s)
	mux.Handle("/cosmo-auth-server"+path, s.handlerArgsMiddleware(http.StripPrefix("/cosmo-auth-server", authServerHandler)))

	// setup serving static files
	mux.Handle("/cosmo-auth-server/",
		http.StripPrefix("/cosmo-auth-server", http.FileServer(http.Dir(s.StaticFileDir))))

	// setup middlewares for all routers to use HTTPRequestLogger and TimeoutHandler.
	// deadline of the Timeout handler takes precedence over any subsequent deadlines
	reqLogr := NewHTTPRequestLogger(s.Log)
	s.http.Handler = reqLogr.Middleware(s.timeoutHandler(mux))

	return mux
}

func (s *Server) setupSessionStore() {
	s.sessionStore = session.NewStore([]byte(s.CookieHashKey), []byte(s.CookieBlockKey), s.sessionOption())
}

func (s *Server) sessionOption() *http.Cookie {
	return &http.Cookie{
		Name:     s.SessionName,
		MaxAge:   s.MaxAgeSeconds,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Domain:   s.CookieDomain,
	}
}

func (s *Server) timeoutHandler(next http.Handler) http.Handler {
	return http.TimeoutHandler(next, s.ResponseTimeout, "")
}

// Start run server
func (s *Server) Start(ctx context.Context) error {

	s.setupRouter()
	s.setupSessionStore()

	go func() {
		<-ctx.Done()
		s.Log.Info("shutdown server")
		s.shutdown()
	}()

	if s.Insecure {
		s.Log.Info("WARNING: start insecure server")
		return s.http.ListenAndServe()

	} else {
		s.Log.Info("start server")
		return s.http.ListenAndServeTLS(s.TLSCertPath, s.TLSPrivateKeyPath)
	}
}

func (s *Server) shutdown() error {
	gracefulShutdownCtx, cancel := context.WithTimeout(context.Background(), s.GracefulShutdownDur)
	defer cancel()
	return s.http.Shutdown(gracefulShutdownCtx)
}
