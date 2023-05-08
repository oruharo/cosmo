package authserver

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/cosmo-workspace/cosmo/pkg/clog"
)

func TestStoreStatusResponseWriter_StatusCode(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		statusCode     int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &StoreStatusResponseWriter{
				ResponseWriter: tt.fields.ResponseWriter,
				statusCode:     tt.fields.statusCode,
			}
			if got := w.StatusCode(); got != tt.want {
				t.Errorf("StoreStatusResponseWriter.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreStatusResponseWriter_WriteHeader(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		statusCode     int
	}
	type args struct {
		statusCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &StoreStatusResponseWriter{
				ResponseWriter: tt.fields.ResponseWriter,
				statusCode:     tt.fields.statusCode,
			}
			w.WriteHeader(tt.args.statusCode)
		})
	}
}

func TestStoreStatusResponseWriter_StatusString(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		statusCode     int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &StoreStatusResponseWriter{
				ResponseWriter: tt.fields.ResponseWriter,
				statusCode:     tt.fields.statusCode,
			}
			if got := w.StatusString(); got != tt.want {
				t.Errorf("StoreStatusResponseWriter.StatusString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHTTPRequestLogger(t *testing.T) {
	type args struct {
		logr *clog.Logger
	}
	tests := []struct {
		name string
		args args
		want HTTPRequestLogger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPRequestLogger(tt.args.logr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPRequestLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPRequestLogger_Middleware(t *testing.T) {
	type fields struct {
		Logger *clog.Logger
	}
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := HTTPRequestLogger{
				Logger: tt.fields.Logger,
			}
			if got := l.Middleware(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPRequestLogger.Middleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrResponse(t *testing.T) {
	type args struct {
		log *clog.Logger
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ErrResponse(tt.args.log, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("ErrResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
