package webhooks

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cosmov1alpha1 "github.com/cosmo-workspace/cosmo/api/core/v1alpha1"
	wsv1alpha1 "github.com/cosmo-workspace/cosmo/api/workspace/v1alpha1"
	"github.com/cosmo-workspace/cosmo/pkg/kosmo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("User webhook", func() {
	normalUserAddon := cosmov1alpha1.Template{
		ObjectMeta: metav1.ObjectMeta{
			Name: "normal-user-addon-test",
			Labels: map[string]string{
				cosmov1alpha1.LabelKeyTemplateType: wsv1alpha1.TemplateTypeUserAddon,
			},
		},
	}

	defaultUserAddon := cosmov1alpha1.Template{
		ObjectMeta: metav1.ObjectMeta{
			Name: "default-user-addon-test",
			Labels: map[string]string{
				cosmov1alpha1.LabelKeyTemplateType: wsv1alpha1.TemplateTypeUserAddon,
			},
			Annotations: map[string]string{
				wsv1alpha1.TemplateAnnKeyDefaultUserAddon: "true",
			},
		},
	}

	notUserAddon := cosmov1alpha1.Template{
		ObjectMeta: metav1.ObjectMeta{
			Name: "notUserAddonTest",
		},
	}

	Context("when creating user with existing addon", func() {
		It("should pass", func() {
			ctx := context.Background()

			var err error
			err = k8sClient.Create(ctx, &normalUserAddon)
			Expect(err).ShouldNot(HaveOccurred())

			err = k8sClient.Create(ctx, &defaultUserAddon)
			Expect(err).ShouldNot(HaveOccurred())

			user := wsv1alpha1.User{}
			user.SetName("testuser1")
			user.Spec = wsv1alpha1.UserSpec{
				AuthType: wsv1alpha1.UserAuthTypeKosmoSecert,
				Addons: []wsv1alpha1.UserAddon{
					{Template: cosmov1alpha1.TemplateRef{Name: defaultUserAddon.GetName()}},
					{Template: cosmov1alpha1.TemplateRef{Name: normalUserAddon.GetName()}},
				},
			}

			err = k8sClient.Create(ctx, &user)
			Expect(err).ShouldNot(HaveOccurred())

			var createdUser wsv1alpha1.User
			Eventually(func() error {
				err := k8sClient.Get(ctx, client.ObjectKey{Name: user.GetName()}, &createdUser)
				if err != nil {
					return err
				}
				return nil
			}, time.Second*10).Should(Succeed())

			// eq := kosmo.LooseDeepEqual(user.DeepCopy(), createdUser.DeepCopy(), kosmo.WithPrintDiff())
			// Expect(eq).Should(BeTrue())
		})
	})

	Context("when creating user with no default addon", func() {
		It("should pass with defaulting", func() {
			ctx := context.Background()

			user := wsv1alpha1.User{
				TypeMeta: metav1.TypeMeta{
					Kind:       "User",
					APIVersion: wsv1alpha1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "testuser2",
				},
				Spec: wsv1alpha1.UserSpec{
					AuthType: wsv1alpha1.UserAuthTypeKosmoSecert,
				},
			}

			expectedUser := user.DeepCopy()
			expectedUser.Spec.Addons = []wsv1alpha1.UserAddon{
				{Template: cosmov1alpha1.TemplateRef{Name: defaultUserAddon.GetName()}},
			}

			err := k8sClient.Create(ctx, &user)
			Expect(err).ShouldNot(HaveOccurred())

			var createdUser wsv1alpha1.User
			Eventually(func() error {
				err := k8sClient.Get(ctx, client.ObjectKey{Name: user.GetName()}, &createdUser)
				if err != nil {
					return err
				}
				return nil
			}, time.Second*10).Should(Succeed())

			expectedUser.ObjectMeta = createdUser.ObjectMeta

			eq := kosmo.LooseDeepEqual(expectedUser, createdUser.DeepCopy(), kosmo.WithPrintDiff())
			Expect(eq).Should(BeTrue())
		})
	})

	Context("when creating user with non-existing addon", func() {
		It("should deny", func() {
			ctx := context.Background()

			user := wsv1alpha1.User{}
			user.SetName("testuser3")
			user.Spec = wsv1alpha1.UserSpec{
				AuthType: wsv1alpha1.UserAuthTypeKosmoSecert,
				Addons: []wsv1alpha1.UserAddon{
					{Template: cosmov1alpha1.TemplateRef{Name: defaultUserAddon.GetName()}},
					{Template: cosmov1alpha1.TemplateRef{Name: "notfound"}},
				},
			}
			err := k8sClient.Create(ctx, &user)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("when creating user with addon template which is not labeled as addon", func() {
		It("should deny", func() {
			ctx := context.Background()

			user := wsv1alpha1.User{}
			user.SetName("testuser4")
			user.Spec = wsv1alpha1.UserSpec{
				AuthType: wsv1alpha1.UserAuthTypeKosmoSecert,
				Addons: []wsv1alpha1.UserAddon{
					{Template: cosmov1alpha1.TemplateRef{Name: defaultUserAddon.GetName()}},
					{Template: cosmov1alpha1.TemplateRef{Name: notUserAddon.GetName()}},
				},
			}
			err := k8sClient.Create(ctx, &user)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("when creating user with no auth type", func() {
		It("should pass with defaulting", func() {
			ctx := context.Background()

			user := wsv1alpha1.User{
				TypeMeta: metav1.TypeMeta{
					Kind:       "User",
					APIVersion: wsv1alpha1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "testuser5",
				},
				Spec: wsv1alpha1.UserSpec{
					Addons: []wsv1alpha1.UserAddon{
						{Template: cosmov1alpha1.TemplateRef{Name: defaultUserAddon.GetName()}},
					},
				},
			}

			expectedUser := user.DeepCopy()
			expectedUser.Spec.AuthType = wsv1alpha1.UserAuthTypeKosmoSecert

			err := k8sClient.Create(ctx, &user)
			Expect(err).ShouldNot(HaveOccurred())

			var createdUser wsv1alpha1.User
			Eventually(func() error {
				err := k8sClient.Get(ctx, client.ObjectKey{Name: user.GetName()}, &createdUser)
				if err != nil {
					return err
				}
				return nil
			}, time.Second*10).Should(Succeed())

			expectedUser.ObjectMeta = createdUser.ObjectMeta

			eq := kosmo.LooseDeepEqual(expectedUser, createdUser.DeepCopy())
			Expect(eq).Should(BeTrue())
		})
	})

	Context("when creating user with innvalid auth type", func() {
		It("should deny", func() {
			ctx := context.Background()

			user := wsv1alpha1.User{}
			user.SetName("testuser6")
			user.Spec = wsv1alpha1.UserSpec{
				AuthType: "invalid",
				Addons: []wsv1alpha1.UserAddon{
					{Template: cosmov1alpha1.TemplateRef{Name: defaultUserAddon.GetName()}},
				},
			}
			err := k8sClient.Create(ctx, &user)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("when creating user with invalid role", func() {
		It("should deny", func() {
			ctx := context.Background()

			user := wsv1alpha1.User{}
			user.SetName("testuser7")
			user.Spec = wsv1alpha1.UserSpec{
				Role:     "invalid",
				AuthType: wsv1alpha1.UserAuthTypeKosmoSecert,
				Addons: []wsv1alpha1.UserAddon{
					{Template: cosmov1alpha1.TemplateRef{Name: defaultUserAddon.GetName()}},
				},
			}
			err := k8sClient.Create(ctx, &user)
			Expect(err).Should(HaveOccurred())
		})
	})
})
