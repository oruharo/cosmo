/*
Copyright © 2023 NAME HERE cosmo-workspace
*/
package authserver

import (
	"reflect"
	"testing"

	"github.com/cosmo-workspace/cosmo/pkg/clog"
	"github.com/spf13/cobra"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func TestNewRootCmd(t *testing.T) {
	type args struct {
		o *options
	}
	tests := []struct {
		name string
		args args
		want *cobra.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRootCmd(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRootCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_PreRunE(t *testing.T) {
	type fields struct {
		KubeConfigPath          string
		KubeContext             string
		ZapOpts                 zap.Options
		Logr                    *clog.Logger
		StaticFileDir           string
		CookieDomain            string
		ResponseTimeoutSeconds  int64
		GracefulShutdownSeconds int64
		TLSPrivateKeyPath       string
		TLSCertPath             string
		Insecure                bool
		ServerPort              int
		MaxAgeMinutes           int
	}
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &options{
				KubeConfigPath:          tt.fields.KubeConfigPath,
				KubeContext:             tt.fields.KubeContext,
				ZapOpts:                 tt.fields.ZapOpts,
				Logr:                    tt.fields.Logr,
				StaticFileDir:           tt.fields.StaticFileDir,
				CookieDomain:            tt.fields.CookieDomain,
				ResponseTimeoutSeconds:  tt.fields.ResponseTimeoutSeconds,
				GracefulShutdownSeconds: tt.fields.GracefulShutdownSeconds,
				TLSPrivateKeyPath:       tt.fields.TLSPrivateKeyPath,
				TLSCertPath:             tt.fields.TLSCertPath,
				Insecure:                tt.fields.Insecure,
				ServerPort:              tt.fields.ServerPort,
				MaxAgeMinutes:           tt.fields.MaxAgeMinutes,
			}
			if err := o.PreRunE(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("options.PreRunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_options_Validate(t *testing.T) {
	type fields struct {
		KubeConfigPath          string
		KubeContext             string
		ZapOpts                 zap.Options
		Logr                    *clog.Logger
		StaticFileDir           string
		CookieDomain            string
		ResponseTimeoutSeconds  int64
		GracefulShutdownSeconds int64
		TLSPrivateKeyPath       string
		TLSCertPath             string
		Insecure                bool
		ServerPort              int
		MaxAgeMinutes           int
	}
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &options{
				KubeConfigPath:          tt.fields.KubeConfigPath,
				KubeContext:             tt.fields.KubeContext,
				ZapOpts:                 tt.fields.ZapOpts,
				Logr:                    tt.fields.Logr,
				StaticFileDir:           tt.fields.StaticFileDir,
				CookieDomain:            tt.fields.CookieDomain,
				ResponseTimeoutSeconds:  tt.fields.ResponseTimeoutSeconds,
				GracefulShutdownSeconds: tt.fields.GracefulShutdownSeconds,
				TLSPrivateKeyPath:       tt.fields.TLSPrivateKeyPath,
				TLSCertPath:             tt.fields.TLSCertPath,
				Insecure:                tt.fields.Insecure,
				ServerPort:              tt.fields.ServerPort,
				MaxAgeMinutes:           tt.fields.MaxAgeMinutes,
			}
			if err := o.Validate(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("options.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_options_Complete(t *testing.T) {
	type fields struct {
		KubeConfigPath          string
		KubeContext             string
		ZapOpts                 zap.Options
		Logr                    *clog.Logger
		StaticFileDir           string
		CookieDomain            string
		ResponseTimeoutSeconds  int64
		GracefulShutdownSeconds int64
		TLSPrivateKeyPath       string
		TLSCertPath             string
		Insecure                bool
		ServerPort              int
		MaxAgeMinutes           int
	}
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &options{
				KubeConfigPath:          tt.fields.KubeConfigPath,
				KubeContext:             tt.fields.KubeContext,
				ZapOpts:                 tt.fields.ZapOpts,
				Logr:                    tt.fields.Logr,
				StaticFileDir:           tt.fields.StaticFileDir,
				CookieDomain:            tt.fields.CookieDomain,
				ResponseTimeoutSeconds:  tt.fields.ResponseTimeoutSeconds,
				GracefulShutdownSeconds: tt.fields.GracefulShutdownSeconds,
				TLSPrivateKeyPath:       tt.fields.TLSPrivateKeyPath,
				TLSCertPath:             tt.fields.TLSCertPath,
				Insecure:                tt.fields.Insecure,
				ServerPort:              tt.fields.ServerPort,
				MaxAgeMinutes:           tt.fields.MaxAgeMinutes,
			}
			if err := o.Complete(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("options.Complete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_options_RunE(t *testing.T) {
	type fields struct {
		KubeConfigPath          string
		KubeContext             string
		ZapOpts                 zap.Options
		Logr                    *clog.Logger
		StaticFileDir           string
		CookieDomain            string
		ResponseTimeoutSeconds  int64
		GracefulShutdownSeconds int64
		TLSPrivateKeyPath       string
		TLSCertPath             string
		Insecure                bool
		ServerPort              int
		MaxAgeMinutes           int
	}
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &options{
				KubeConfigPath:          tt.fields.KubeConfigPath,
				KubeContext:             tt.fields.KubeContext,
				ZapOpts:                 tt.fields.ZapOpts,
				Logr:                    tt.fields.Logr,
				StaticFileDir:           tt.fields.StaticFileDir,
				CookieDomain:            tt.fields.CookieDomain,
				ResponseTimeoutSeconds:  tt.fields.ResponseTimeoutSeconds,
				GracefulShutdownSeconds: tt.fields.GracefulShutdownSeconds,
				TLSPrivateKeyPath:       tt.fields.TLSPrivateKeyPath,
				TLSCertPath:             tt.fields.TLSCertPath,
				Insecure:                tt.fields.Insecure,
				ServerPort:              tt.fields.ServerPort,
				MaxAgeMinutes:           tt.fields.MaxAgeMinutes,
			}
			if err := o.RunE(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("options.RunE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute()
		})
	}
}

func Test_printVersion(t *testing.T) {
	type args struct {
		cmd *cobra.Command
		o   *options
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printVersion(tt.args.cmd, tt.args.o)
		})
	}
}

func Test_printOptions(t *testing.T) {
	type args struct {
		o *options
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printOptions(tt.args.o)
		})
	}
}
