package kosmo

import (
	"context"

	"k8s.io/apimachinery/pkg/types"

	cosmov1alpha1 "github.com/cosmo-workspace/cosmo/api/v1alpha1"
	"github.com/cosmo-workspace/cosmo/pkg/clog"
	"github.com/cosmo-workspace/cosmo/pkg/kubeutil"
	"github.com/cosmo-workspace/cosmo/pkg/role"
)

func filterTemplates(ctx context.Context, tmpls []cosmov1alpha1.TemplateObject, roles []cosmov1alpha1.UserRole) []cosmov1alpha1.TemplateObject {
	filteredTmpls := make([]cosmov1alpha1.TemplateObject, 0, len(tmpls))
	for _, v := range tmpls {
		if role.IsAllowedToUseTemplate(ctx, v, roles) {
			filteredTmpls = append(filteredTmpls, v)
		}
	}
	return filteredTmpls
}

func (c *Client) ListWorkspaceTemplates(ctx context.Context, roles []cosmov1alpha1.UserRole) ([]cosmov1alpha1.TemplateObject, error) {
	log := clog.FromContext(ctx).WithCaller()
	if tmpls, err := kubeutil.ListTemplateObjectsByType(ctx, c, []string{cosmov1alpha1.TemplateLabelEnumTypeWorkspace}); err != nil {
		log.Error(err, "failed to list WorkspaceTemplates")
		return nil, NewInternalServerError("failed to list WorkspaceTemplates", err)
	} else {
		return filterTemplates(ctx, tmpls, roles), nil
	}
}

func (c *Client) ListUserAddonTemplates(ctx context.Context, roles []cosmov1alpha1.UserRole) ([]cosmov1alpha1.TemplateObject, error) {
	log := clog.FromContext(ctx).WithCaller()
	if tmpls, err := kubeutil.ListTemplateObjectsByType(ctx, c, []string{cosmov1alpha1.TemplateLabelEnumTypeUserAddon}); err != nil {
		log.Error(err, "failed to list UserAddon Templates")
		return nil, NewInternalServerError("failed to list UserAddon Templates", err)
	} else {
		return filterTemplates(ctx, tmpls, roles), nil
	}
}

func (c *Client) GetTemplate(ctx context.Context, tmplName string) (*cosmov1alpha1.Template, error) {
	tmpl := cosmov1alpha1.Template{}

	key := types.NamespacedName{
		Name: tmplName,
	}

	if err := c.Get(ctx, key, &tmpl); err != nil {
		return nil, err
	}
	return &tmpl, nil
}
