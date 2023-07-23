package role

import (
	"context"
	"testing"

	cosmov1alpha1 "github.com/cosmo-workspace/cosmo/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_isAllowedToUseTemplate(t *testing.T) {
	type args struct {
		tmpl  cosmov1alpha1.TemplateObject
		roles []cosmov1alpha1.UserRole
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no annotations, all roles are allowed",
			args: args{
				tmpl: &cosmov1alpha1.Template{
					ObjectMeta: metav1.ObjectMeta{
						Name: "hogwarts-common",
					},
				},
				roles: []cosmov1alpha1.UserRole{
					{Name: "gryffindor-developer"},
				},
			},
			want: true,
		},
		{
			name: "forbidden if role is matched to forbidden role",
			args: args{
				tmpl: &cosmov1alpha1.Template{
					ObjectMeta: metav1.ObjectMeta{
						Name: "sword-of-gryffindor",
						Annotations: map[string]string{
							cosmov1alpha1.TemplateAnnKeyForbiddenUserRoles: "slytherin",
						},
					},
				},
				roles: []cosmov1alpha1.UserRole{
					{Name: "slytherin"},
				},
			},
			want: false,
		},
		{
			name: "forbidden if role is not matched to allowed role",
			args: args{
				tmpl: &cosmov1alpha1.Template{
					ObjectMeta: metav1.ObjectMeta{
						Name: "sword-of-gryffindor",
						Annotations: map[string]string{
							cosmov1alpha1.TemplateAnnKeyUserRoles: "gryffindor",
						},
					},
				},
				roles: []cosmov1alpha1.UserRole{
					{Name: "slytherin"},
				},
			},
			want: false,
		},
		{
			name: "forbidden if role is matched to allowed role but also matched to forbidden role",
			args: args{
				tmpl: &cosmov1alpha1.Template{
					ObjectMeta: metav1.ObjectMeta{
						Name: "sword-of-gryffindor",
						Annotations: map[string]string{
							cosmov1alpha1.TemplateAnnKeyForbiddenUserRoles: "slytherin",
							cosmov1alpha1.TemplateAnnKeyUserRoles:          "gryffindor",
						},
					},
				},
				roles: []cosmov1alpha1.UserRole{
					{Name: "slytherin"},
					{Name: "gryffindor"},
				},
			},
			want: false,
		},
		{
			name: "allowed if wildcard match for allowed role",
			args: args{
				tmpl: &cosmov1alpha1.Template{
					ObjectMeta: metav1.ObjectMeta{
						Name: "sword-of-gryffindor",
						Annotations: map[string]string{
							cosmov1alpha1.TemplateAnnKeyUserRoles: "gryffindor-*",
						},
					},
				},
				roles: []cosmov1alpha1.UserRole{
					{Name: "gryffindor-developer"},
				},
			},
			want: true,
		},
		{
			name: "forbidden if wildcard match for forbidden role",
			args: args{
				tmpl: &cosmov1alpha1.Template{
					ObjectMeta: metav1.ObjectMeta{
						Name: "sword-of-gryffindor",
						Annotations: map[string]string{
							cosmov1alpha1.TemplateAnnKeyUserRoles: "sly*",
						},
					},
				},
				roles: []cosmov1alpha1.UserRole{
					{Name: "slytherin"},
				},
			},
			want: true,
		},
		{
			name: "forbidden if allowed role wildcard not match",
			args: args{
				tmpl: &cosmov1alpha1.Template{
					ObjectMeta: metav1.ObjectMeta{
						Name: "sword-of-gryffindor",
						Annotations: map[string]string{
							cosmov1alpha1.TemplateAnnKeyUserRoles: "gryffindor-*",
						},
					},
				},
				roles: []cosmov1alpha1.UserRole{
					{Name: "gryffindor"},
				},
			},
			want: false,
		},
		{
			name: "forbidden if both allowed role wildcard and forbidden role matches",
			args: args{
				tmpl: &cosmov1alpha1.Template{
					ObjectMeta: metav1.ObjectMeta{
						Name: "sword-of-gryffindor",
						Annotations: map[string]string{
							cosmov1alpha1.TemplateAnnKeyUserRoles:          "gryffindor-*",
							cosmov1alpha1.TemplateAnnKeyForbiddenUserRoles: "gryffindor-faker",
						},
					},
				},
				roles: []cosmov1alpha1.UserRole{
					{Name: "gryffindor-faker"},
				},
			},
			want: false,
		},
		{
			name: "forbidden if both allowed role wildcard and forbidden wildcard matches",
			args: args{
				tmpl: &cosmov1alpha1.Template{
					ObjectMeta: metav1.ObjectMeta{
						Name: "sword-of-gryffindor",
						Annotations: map[string]string{
							cosmov1alpha1.TemplateAnnKeyUserRoles:          "gryffindor-*",
							cosmov1alpha1.TemplateAnnKeyForbiddenUserRoles: "gryffindor-f*",
						},
					},
				},
				roles: []cosmov1alpha1.UserRole{
					{Name: "gryffindor-faker"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllowedToUseTemplate(context.TODO(), tt.args.tmpl, tt.args.roles); got != tt.want {
				t.Errorf("isAllowedToUseTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
