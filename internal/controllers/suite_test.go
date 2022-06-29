package controllers

import (
	"context"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	cosmov1alpha1 "github.com/cosmo-workspace/cosmo/api/core/v1alpha1"
	wsv1alpha1 "github.com/cosmo-workspace/cosmo/api/workspace/v1alpha1"

	//+kubebuilder:scaffold:imports

	"github.com/cosmo-workspace/cosmo/pkg/kosmo"
)

const (
	controllerFieldManager string = "cosmo-instance-controller"
)

const (
	instController        string = "cosmo-instance-controller"
	clusterInstController string = "cosmo-cluster-instance-controller"
	tmplController        string = "cosmo-template-controller"
	clusterTmplController string = "cosmo-cluster-template-controller"
	userController        string = "cosmo-user-controller"
	wsController          string = "cosmo-workspace-controller"
	wsStatController      string = "cosmo-workspace-status-controller"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var (
	cfg       *rest.Config
	k8sClient kosmo.Client
	testEnv   *envtest.Environment
	ctx       context.Context
	cancel    context.CancelFunc
)

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	ctx, cancel = context.WithCancel(ctrl.SetupSignalHandler())

	By("bootstrapping test environment")

	testEnv = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "..", "config", "crd", "bases")},
		ErrorIfCRDPathMissing: true,
	}

	var err error
	cfg, err = testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	err = cosmov1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = wsv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	//+kubebuilder:scaffold:scheme

	c, err := client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())

	k8sClient = kosmo.NewClient(c)

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme:             scheme.Scheme,
		MetricsBindAddress: "0",
	})
	Expect(err).NotTo(HaveOccurred())

	err = (&InstanceReconciler{
		Client:   kosmo.NewClient(mgr.GetClient()),
		Scheme:   mgr.GetScheme(),
		Recorder: mgr.GetEventRecorderFor(instController),
	}).SetupWithManager(mgr, controllerFieldManager)
	Expect(err).NotTo(HaveOccurred())

	err = (&ClusterInstanceReconciler{
		Client:   kosmo.NewClient(mgr.GetClient()),
		Scheme:   mgr.GetScheme(),
		Recorder: mgr.GetEventRecorderFor(instController),
	}).SetupWithManager(mgr, controllerFieldManager)
	Expect(err).NotTo(HaveOccurred())

	err = (&TemplateReconciler{
		Client: kosmo.NewClient(mgr.GetClient()),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr)
	Expect(err).NotTo(HaveOccurred())

	err = (&ClusterTemplateReconciler{
		Client: kosmo.NewClient(mgr.GetClient()),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr)
	Expect(err).NotTo(HaveOccurred())

	err = (&WorkspaceReconciler{
		Client:   kosmo.NewClient(mgr.GetClient()),
		Scheme:   mgr.GetScheme(),
		Recorder: mgr.GetEventRecorderFor(wsController),
	}).SetupWithManager(mgr)
	Expect(err).NotTo(HaveOccurred())

	err = (&WorkspaceStatusReconciler{
		Client:         kosmo.NewClient(mgr.GetClient()),
		Scheme:         mgr.GetScheme(),
		Recorder:       mgr.GetEventRecorderFor(wsStatController),
		DefaultURLBase: "",
	}).SetupWithManager(mgr)
	Expect(err).NotTo(HaveOccurred())

	err = (&UserReconciler{
		Client:   kosmo.NewClient(mgr.GetClient()),
		Scheme:   mgr.GetScheme(),
		Recorder: mgr.GetEventRecorderFor(userController),
	}).SetupWithManager(mgr)
	Expect(err).NotTo(HaveOccurred())

	go func() {
		defer GinkgoRecover()
		err := mgr.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	}()

	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())

})

var _ = AfterSuite(func() {
	cancel()
	By("tearing down the test environment")
	err := testEnv.Stop()
	Expect(err).NotTo(HaveOccurred())
})
