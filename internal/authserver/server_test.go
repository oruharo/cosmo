package authserver

import (
	"context"
	"net/http"
	"testing"

	"github.com/cosmo-workspace/cosmo/pkg/clog"
	"github.com/gkampitakis/go-snaps/snaps"
	ctrl "sigs.k8s.io/controller-runtime"
)

func TestServer_setupRouter(t *testing.T) {

	tests := []struct {
		name   string
		server *Server
	}{
		{
			server: &Server{
				Log:  clog.NewLogger(ctrl.Log.WithName("authserver")),
				http: &http.Server{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.server
			s.setupRouter()
			snaps.MatchSnapshot(t, s.setupRouter())
			snaps.MatchSnapshot(t, s)
		})
	}
}

func TestServer_setupSessionStore(t *testing.T) {

	tests := []struct {
		name   string
		server *Server
	}{
		{
			server: &Server{
				Log:  clog.NewLogger(ctrl.Log.WithName("authserver")),
				http: &http.Server{},
			},
		},
		{
			server: &Server{
				Log:            clog.NewLogger(ctrl.Log.WithName("authserver")),
				CookieHashKey:  "abcde",
				CookieBlockKey: "efghi",
				http:           &http.Server{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.server
			s.setupSessionStore()
			snaps.MatchSnapshot(t, s)
		})
	}
}

func TestServer_timeoutHandler(t *testing.T) {
	tests := []struct {
		name   string
		server *Server
	}{
		{
			server: &Server{
				Log:  clog.NewLogger(ctrl.Log.WithName("authserver")),
				http: &http.Server{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.server
			next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
			handler := s.timeoutHandler(next)
			snaps.MatchSnapshot(t, handler)
		})
	}
}

func TestServer_Start(t *testing.T) {
	tests := []struct {
		name   string
		server *Server
	}{
		{
			server: &Server{
				Log:  clog.NewLogger(ctrl.Log.WithName("authserver")),
				http: &http.Server{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.server
			ctx := context.Background()
			s.Start(ctx)
			snaps.MatchSnapshot(t, s)
			s.shutdown()
			s.Start(ctx)
		})
	}
}
