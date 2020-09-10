package controllers

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Ingress controller", func() {

	// Define utility constants for object names and testing timeouts/durations and intervals.
	const (
		name       = "test"
		namespace  = "default"
		hostMask   = "test.0.0.0.0.xip.io"
		hostWithIP = "test.1.2.3.4.xip.io"
		ip         = "1.2.3.4"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("when an Ingress has hosts matching IP pattern", func() {
		It("should replace that pattern with the Load Balancer IP address", func() {
			By("Creating a new Ingress")

			ctx := context.Background()
			key := types.NamespacedName{
				Name:      name,
				Namespace: namespace,
			}
			ingress := &extensionsv1beta1.Ingress{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
				},
				Spec: extensionsv1beta1.IngressSpec{
					Rules: []extensionsv1beta1.IngressRule{
						{
							Host: hostMask,
						},
					},
					TLS: []extensionsv1beta1.IngressTLS{
						{
							Hosts: []string{
								hostMask,
							},
							SecretName: "secret",
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, ingress)).Should(Succeed())

			found := &extensionsv1beta1.Ingress{}
			Eventually(func() bool {
				if err := k8sClient.Get(ctx, key, found); err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())
			Consistently(func() ([]string, error) {
				if err := k8sClient.Get(ctx, key, found); err != nil {
					return nil, err
				}
				var hosts []string
				for _, rule := range found.Spec.Rules {
					hosts = append(hosts, rule.Host)
				}
				for _, tls := range found.Spec.TLS {
					hosts = append(hosts, tls.Hosts...)
				}
				return hosts, nil
			}, duration, interval).Should(ConsistOf(hostMask, hostMask))

			By("Assigning a Load Balance IP address")

			found.Status.LoadBalancer.Ingress = append(found.Status.LoadBalancer.Ingress, corev1.LoadBalancerIngress{
				IP: ip,
			})
			Expect(k8sClient.Status().Update(ctx, found)).Should(Succeed())

			found = &extensionsv1beta1.Ingress{}
			Eventually(func() []corev1.LoadBalancerIngress {
				_ = k8sClient.Get(ctx, key, found)
				return found.Status.LoadBalancer.Ingress
			}, timeout, interval).ShouldNot(BeEmpty())
			Expect(found.Status.LoadBalancer.Ingress[0].IP).Should(Equal(ip))

			found = &extensionsv1beta1.Ingress{}
			Eventually(func() ([]string, error) {
				if err := k8sClient.Get(ctx, key, found); err != nil {
					return nil, err
				}
				var hosts []string
				for _, rule := range found.Spec.Rules {
					hosts = append(hosts, rule.Host)
				}
				for _, tls := range found.Spec.TLS {
					hosts = append(hosts, tls.Hosts...)
				}
				return hosts, nil
			}, timeout, interval).Should(ConsistOf(hostWithIP, hostWithIP))
		})
	})
})
