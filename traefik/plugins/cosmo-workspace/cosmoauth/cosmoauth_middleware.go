package cosmoauth

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"strings"
	"time"

	"github.com/cosmo-workspace/cosmo/pkg/auth/session"
	"github.com/gorilla/sessions"
)

// nolint
var (
	LoggerDEBUG = log.New(io.Discard, "DEBUG: cosmo-auth: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerINFO  = log.New(io.Discard, "INFO: cosmo-auth: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerERROR = log.New(io.Discard, "ERROR: cosmo-auth: ", log.Ldate|log.Ltime|log.Lshortfile)
)

// Config the plugin configuration.
type Config struct {
	LogLevel          string `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
	CookieSessionName string `json:"cookieSessionName,omitempty" yaml:"cookieSessionName,omitempty"`
	CookieDomain      string `json:"cookieDomain,omitempty" yaml:"cookieDomain,omitempty"`
	CookieHashKey     string `json:"cookieHashKey,omitempty" yaml:"cookieHashKey,omitempty"`
	CookieBlockKey    string `json:"cookieBlockKey,omitempty" yaml:"cookieBlockKey,omitempty"`
	SignInUrl         string `json:"signInUrl,omitempty" yaml:"signInUrl,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		LogLevel:          "INFO",
		CookieSessionName: "",
		CookieDomain:      "",
		CookieHashKey:     "----+----X----+----X----+----X----+----X----+----X----+----X----",
		CookieBlockKey:    "----+----X----+----X----+----X--",
		SignInUrl:         "",
	}
}

type CosmoAuth struct {
	config       *Config
	next         http.Handler
	name         string
	RedirectPath string

	SessionStore sessions.Store
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	SetLogger(config.LogLevel)
	LoggerINFO.Printf("Starting %s Middleware...", name)

	conf := &Config{
		LogLevel:          os.ExpandEnv(config.LogLevel),
		CookieSessionName: os.ExpandEnv(config.CookieSessionName),
		CookieDomain:      os.ExpandEnv(config.CookieDomain),
		CookieHashKey:     os.ExpandEnv(config.CookieHashKey),
		CookieBlockKey:    os.ExpandEnv(config.CookieBlockKey),
		SignInUrl:         os.ExpandEnv(config.SignInUrl),
	}

	p := &CosmoAuth{
		config:       conf,
		next:         next,
		name:         name,
		SessionStore: sessions.NewCookieStore([]byte(conf.CookieHashKey), []byte(conf.CookieBlockKey)),
	}

	return p, nil
}

func (p *CosmoAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Bypass manifest.json not to check session. By default, manifest.json is requested without cookie.
	// https://developer.mozilla.org/en-US/docs/Web/Manifest
	if strings.Contains(strings.ToLower(r.URL.Path), "/manifest.json") {
		p.next.ServeHTTP(w, r)
		return
	}
	ses, err := p.SessionStore.Get(r, p.config.CookieSessionName)
	if ses == nil || err != nil {
		LoggerERROR.Printf("failed to get session from store. err=%s", err)
		p.redirectToLoginPage(w, r)
		return
	}
	if ses.IsNew {
		LoggerINFO.Println("not authorized")
		p.redirectToLoginPage(w, r)
		return
	}

	sesInfo := session.Get(ses)
	LoggerDEBUG.Print("get session.", " userName=", sesInfo.UserName, " deadline=", sesInfo.Deadline)

	// check user name is owner's
	userName := r.Header.Get("X-Cosmo-UserName")
	// TODO: improvement
	if userName != "" && sesInfo.UserName != userName {
		LoggerINFO.Print("invalid authorization.", " storedUserName=", sesInfo.UserName, " ownerName=", userName)
		p.redirectToLoginPage(w, r)
		return
	}

	// set deadline on request if enabled
	ctx := r.Context()
	if sesInfo.Deadline > 0 {
		deadline := time.Unix(sesInfo.Deadline, 0)
		LoggerDEBUG.Print("set deadline at ", deadline)

		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}

	LoggerDEBUG.Print("authorized.", " path=", r.URL.Path)
	p.next.ServeHTTP(w, r.WithContext(ctx))
	w.Header().Set("X-Cosmo-UserName", sesInfo.UserName)
}

func (p *CosmoAuth) redirectToLoginPage(w http.ResponseWriter, r *http.Request) {

	bodyFormat := `
<!DOCTYPE html>
<html lang="en">
	<head>
		<title>COSMO Auth redirector</title>
		<script type="module">
			const originalUrl = encodeURIComponent(window.location.href);
			const signInUrl = "%s" + "?redirect_to=" + originalUrl;
			window.location.href = signInUrl;
			console.log(signInUrl)
		</script>
	</head>
	<body>redirect to Sign In page</body>
</html>
`
	// body := fmt.Sprintf(bodyFormat, p.config.SignInUrl)
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, bodyFormat, p.config.SignInUrl)
	// w.Write([]byte(body))
}

func SetLogger(level string) {
	switch level {
	case "ERROR":
		LoggerERROR.SetOutput(os.Stderr)
	case "INFO":
		LoggerERROR.SetOutput(os.Stderr)
		LoggerINFO.SetOutput(os.Stdout)
	case "DEBUG":
		LoggerERROR.SetOutput(os.Stderr)
		LoggerINFO.SetOutput(os.Stdout)
		LoggerDEBUG.SetOutput(os.Stdout)
	default:
		LoggerERROR.SetOutput(os.Stderr)
		LoggerINFO.SetOutput(os.Stdout)
	}
}