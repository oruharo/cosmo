package dashboard

import (
	"context"
	"fmt"
	"reflect"

	connect_go "github.com/bufbuild/connect-go"

	cosmov1alpha1 "github.com/cosmo-workspace/cosmo/api/v1alpha1"
	"github.com/cosmo-workspace/cosmo/pkg/clog"
	"github.com/cosmo-workspace/cosmo/pkg/kosmo"
	"github.com/cosmo-workspace/cosmo/pkg/slices"
	dashv1alpha1 "github.com/cosmo-workspace/cosmo/proto/gen/dashboard/v1alpha1"
)

func (s *Server) UpdateUserDisplayName(ctx context.Context, req *connect_go.Request[dashv1alpha1.UpdateUserDisplayNameRequest]) (*connect_go.Response[dashv1alpha1.UpdateUserDisplayNameResponse], error) {
	log := clog.FromContext(ctx).WithCaller()
	log.Debug().Info("request", "req", req)

	if err := userAuthentication(ctx, req.Msg.UserName); err != nil {
		return nil, ErrResponse(log, err)
	}

	user, err := s.Klient.UpdateUser(ctx, req.Msg.UserName, kosmo.UpdateUserOpts{
		DisplayName: &req.Msg.DisplayName,
		UserRoles:   []string{"-"}})
	if err != nil {
		return nil, ErrResponse(log, err)
	}

	res := &dashv1alpha1.UpdateUserDisplayNameResponse{
		Message: "Successfully updated",
		User:    convertUserToDashv1alpha1User(*user),
	}
	log.Info(res.Message, "username", req.Msg.UserName)
	return connect_go.NewResponse(res), nil
}

func diff(slice1 []string, slice2 []string) []string {
	var diff []string
	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices
		slice1, slice2 = slice2, slice1
	}
	return diff
}

func (s *Server) UpdateUserRole(ctx context.Context, req *connect_go.Request[dashv1alpha1.UpdateUserRoleRequest]) (*connect_go.Response[dashv1alpha1.UpdateUserRoleResponse], error) {
	log := clog.FromContext(ctx).WithCaller()
	log.Debug().Info("request", "req", req)

	currentUser, err := s.Klient.GetUser(ctx, req.Msg.UserName)
	if err != nil {
		return nil, ErrResponse(log, err)
	}

	// group-admin can attach or detach only group-roles
	changingRoles := diff(convertUserRolesToStringSlice(currentUser.Spec.Roles), req.Msg.Roles)
	if err := adminAuthentication(ctx, validateCallerHasAdminForAllRoles(changingRoles)); err != nil {
		return nil, ErrResponse(log, err)
	}

	user, err := s.Klient.UpdateUser(ctx, req.Msg.UserName, kosmo.UpdateUserOpts{UserRoles: req.Msg.Roles})
	if err != nil {
		return nil, ErrResponse(log, err)
	}

	res := &dashv1alpha1.UpdateUserRoleResponse{
		Message: "Successfully updated",
		User:    convertUserToDashv1alpha1User(*user),
	}
	log.Info(res.Message, "username", req.Msg.UserName)
	return connect_go.NewResponse(res), nil
}

func (s *Server) UpdateUserPassword(ctx context.Context, req *connect_go.Request[dashv1alpha1.UpdateUserPasswordRequest]) (*connect_go.Response[dashv1alpha1.UpdateUserPasswordResponse], error) {
	log := clog.FromContext(ctx).WithCaller()
	log.Debug().Info("request", "username", req.Msg.UserName)

	if err := userAuthentication(ctx, req.Msg.UserName); err != nil {
		return nil, ErrResponse(log, err)
	}

	// check current password is valid
	verified, _, err := s.Klient.VerifyPassword(ctx, req.Msg.UserName, []byte(req.Msg.CurrentPassword))
	if err != nil {
		return nil, ErrResponse(log, err)
	}

	if !verified {
		return nil, ErrResponse(log, kosmo.NewForbiddenError("incorrect user or password", nil))
	}

	// Upsert password
	if err := s.Klient.RegisterPassword(ctx, req.Msg.UserName, []byte(req.Msg.NewPassword)); err != nil {
		return nil, ErrResponse(log, err)
	}

	res := &dashv1alpha1.UpdateUserPasswordResponse{
		Message: "Successfully updated",
	}
	log.Info(res.Message, "username", req.Msg.UserName)
	return connect_go.NewResponse(res), nil
}

func (s *Server) UpdateUserAddons(ctx context.Context, req *connect_go.Request[dashv1alpha1.UpdateUserAddonsRequest]) (*connect_go.Response[dashv1alpha1.UpdateUserAddonsResponse], error) {
	log := clog.FromContext(ctx).WithCaller()
	log.Debug().Info("request", "req", req)

	targetUser, err := s.Klient.GetUser(ctx, req.Msg.UserName)
	if err != nil {
		return nil, ErrResponse(log, err)
	}

	caller := callerFromContext(ctx)
	if caller == nil {
		return nil, kosmo.NewInternalServerError("unable get caller", nil)
	}

	availableTemplates, err := s.Klient.ListUserAddonTemplates(ctx, targetUser.Spec.Roles)
	if err != nil {
		return nil, ErrResponse(log, err)
	}
	changingAddons := diffUserAddons(targetUser.Spec.Addons, convertDashv1alpha1UserAddonToUserAddon(req.Msg.Addons))
	for _, addon := range changingAddons {
		if _, ok := find(availableTemplates, addon,
			func(e cosmov1alpha1.TemplateObject, v cosmov1alpha1.UserAddon) bool {
				return e.GetName() == v.Template.Name
			}); !ok {
			return nil, ErrResponse(log, kosmo.NewForbiddenError(fmt.Sprintf("denied to attach '%s' group", addon.Template.Name), nil))
		}
	}

	addons := convertDashv1alpha1UserAddonToUserAddon(req.Msg.Addons)
	user, err := s.Klient.UpdateUser(ctx, req.Msg.UserName, kosmo.UpdateUserOpts{
		UserRoles: []string{"-"},
		Addons:    &addons,
	})
	if err != nil {
		return nil, ErrResponse(log, err)
	}

	res := &dashv1alpha1.UpdateUserAddonsResponse{
		Message: "Successfully updated",
		User:    convertUserToDashv1alpha1User(*user),
	}
	log.Info(res.Message, "username", req.Msg.UserName)
	return connect_go.NewResponse(res), nil
}

func find[T any, U any](elems []T, v U, compareFunc func(T, U) bool) (matchElm T, result bool) {
	for _, e := range elems {
		if compareFunc(e, v) {
			return e, true
		}
	}
	return matchElm, false
}

func diffAny[T any](before, after []T, compareFunc func(T, T) bool) (diff []T) {
	for _, s := range before {
		if _, ok := find(after, s, compareFunc); !ok {
			diff = append(diff, s) // elements contained only in before
		}
	}
	for _, s := range after {
		v, ok := find(before, s, compareFunc)
		if !ok || !reflect.DeepEqual(v, s) {
			diff = append(diff, s) // elements contained only in after + updated elements
		}
	}
	return diff
}

func diffAnyz[T any](before, after []T, compareFunc func(T, T) bool) (diff []T) {

	for _, s := range before {
		ok := slices.ContainsFunc(after, func(v2 T) bool { return compareFunc(s, v2) })
		if !ok {
			diff = append(diff, s) // elements contained only in before
		}
	}
	for _, s := range after {
		elm := slices.IndexFunc(before, func(v2 T) bool { return compareFunc(s, v2) })
		if elm > 0 || !reflect.DeepEqual(before[elm], s) {
			diff = append(diff, s) // elements contained only in after + updated elements
		}
	}
	return diff
}

func diffStrings(before, after []string) (diff []string) {
	return diffAny(before, after, func(a, b string) bool {
		return a == b
	})
}

func diffUserAddons(before, after []cosmov1alpha1.UserAddon) []cosmov1alpha1.UserAddon {
	return diffAny(before, after, func(a, b cosmov1alpha1.UserAddon) bool {
		return a.Template == b.Template
	})
}
