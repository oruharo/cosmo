package wscfg

import (
	"errors"

	cosmov1alpha1 "github.com/cosmo-workspace/cosmo/api/v1alpha1"
)

var (
	ErrNotTypeWorkspace = errors.New("not type workspace")
)

func SetConfigOnTemplateAnnotations(tmpl cosmov1alpha1.TemplateObject, cfg cosmov1alpha1.Config) {
	ann := tmpl.GetAnnotations()
	if ann == nil {
		ann = make(map[string]string)
	}
	ann[cosmov1alpha1.WorkspaceTemplateAnnKeyDeploymentName] = cfg.DeploymentName
	ann[cosmov1alpha1.WorkspaceTemplateAnnKeyServiceName] = cfg.ServiceName
	ann[cosmov1alpha1.WorkspaceTemplateAnnKeyIngressName] = cfg.IngressName
	ann[cosmov1alpha1.WorkspaceTemplateAnnKeyIngressRouteName] = cfg.IngressRouteName
	ann[cosmov1alpha1.WorkspaceTemplateAnnKeyServiceMainPort] = cfg.ServiceMainPortName
	ann[cosmov1alpha1.WorkspaceTemplateAnnKeyURLBase] = cfg.URLBase
	tmpl.SetAnnotations(ann)
}

func ConfigFromTemplateAnnotations(tmpl *cosmov1alpha1.Template) (cfg cosmov1alpha1.Config, err error) {
	// check TemplateType label is for Workspace
	labels := tmpl.GetLabels()
	if labels == nil {
		return cfg, ErrNotTypeWorkspace
	}

	tmplType, ok := labels[cosmov1alpha1.TemplateLabelKeyType]
	if !ok || tmplType != cosmov1alpha1.TemplateLabelEnumTypeWorkspace {
		return cfg, ErrNotTypeWorkspace
	}

	ann := tmpl.GetAnnotations()
	if ann == nil {
		return cfg, nil
	}

	cfg = cosmov1alpha1.Config{
		DeploymentName:      ann[cosmov1alpha1.WorkspaceTemplateAnnKeyDeploymentName],
		ServiceName:         ann[cosmov1alpha1.WorkspaceTemplateAnnKeyServiceName],
		IngressName:         ann[cosmov1alpha1.WorkspaceTemplateAnnKeyIngressName],
		IngressRouteName:    ann[cosmov1alpha1.WorkspaceTemplateAnnKeyIngressRouteName],
		ServiceMainPortName: ann[cosmov1alpha1.WorkspaceTemplateAnnKeyServiceMainPort],
		URLBase:             ann[cosmov1alpha1.WorkspaceTemplateAnnKeyURLBase],
	}
	return cfg, nil
}
