package dashboard

import (
	"reflect"
	"testing"

	cosmov1alpha1 "github.com/cosmo-workspace/cosmo/api/v1alpha1"
)

func Test_diff(t *testing.T) {

	tests := []struct {
		name     string
		before   []string
		after    []string
		wantDiff []string
	}{
		{
			name:     "",
			before:   []string{"aa", "bb", "cc"},
			after:    []string{"bb", "dd"},
			wantDiff: []string{"aa", "cc", "dd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiff := diff(tt.before, tt.after); !reflect.DeepEqual(gotDiff, tt.wantDiff) {
				t.Errorf("diffx() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}

func Test_diffStrings(t *testing.T) {

	tests := []struct {
		name     string
		before   []string
		after    []string
		wantDiff []string
	}{
		{
			name:     "",
			before:   []string{"aa", "bb", "cc"},
			after:    []string{"bb", "dd"},
			wantDiff: []string{"aa", "cc", "dd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiff := diffStrings(tt.before, tt.after); !reflect.DeepEqual(gotDiff, tt.wantDiff) {
				t.Errorf("diffx() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}

func Test_diffUserAddons(t *testing.T) {

	tests := []struct {
		name     string
		before   []cosmov1alpha1.UserAddon
		after    []cosmov1alpha1.UserAddon
		wantDiff []cosmov1alpha1.UserAddon
	}{
		{
			name: "",
			before: []cosmov1alpha1.UserAddon{
				{Template: cosmov1alpha1.UserAddonTemplateRef{Name: "aa", ClusterScoped: false}, Vars: map[string]string{}},
				{Template: cosmov1alpha1.UserAddonTemplateRef{Name: "bb", ClusterScoped: false}, Vars: map[string]string{}},
				{Template: cosmov1alpha1.UserAddonTemplateRef{Name: "cc", ClusterScoped: false}, Vars: map[string]string{}},
			},
			after: []cosmov1alpha1.UserAddon{
				{Template: cosmov1alpha1.UserAddonTemplateRef{Name: "bb", ClusterScoped: false}, Vars: map[string]string{}},
				{Template: cosmov1alpha1.UserAddonTemplateRef{Name: "cc", ClusterScoped: false}, Vars: map[string]string{"xx": "xxx"}},
				{Template: cosmov1alpha1.UserAddonTemplateRef{Name: "dd", ClusterScoped: false}, Vars: map[string]string{}},
			},
			wantDiff: []cosmov1alpha1.UserAddon{
				{Template: cosmov1alpha1.UserAddonTemplateRef{Name: "aa", ClusterScoped: false}, Vars: map[string]string{}},
				{Template: cosmov1alpha1.UserAddonTemplateRef{Name: "cc", ClusterScoped: false}, Vars: map[string]string{"xx": "xxx"}},
				{Template: cosmov1alpha1.UserAddonTemplateRef{Name: "dd", ClusterScoped: false}, Vars: map[string]string{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiff := diffUserAddons(tt.before, tt.after); !reflect.DeepEqual(gotDiff, tt.wantDiff) {
				t.Errorf("diffx() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}
