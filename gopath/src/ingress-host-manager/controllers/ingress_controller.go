/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-logr/logr"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	hostIpPattern          = ".0.0.0.0."
	ingressClassAnnotation = "kubernetes.io/ingress.class"
	ingressClassSuffix     = "-ihm"
)

// IngressReconciler reconciles a Ingress object
type IngressReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=extensions,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=extensions,resources=ingresses/status,verbs=get;update;patch

func (r *IngressReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	//_ = context.Background()
	ctx := context.TODO()
	logger := r.Log.WithValues("ingress", req.NamespacedName)

	// Fetch the Ingress instance
	instance := &extensionsv1beta1.Ingress{}
	err := r.Client.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	ingressClass := instance.Annotations[ingressClassAnnotation]
	if strings.HasSuffix(ingressClass, ingressClassSuffix) {
		ingress := newIngress(instance)

		found := &extensionsv1beta1.Ingress{}
		err := r.Client.Get(ctx, types.NamespacedName{Namespace: ingress.Namespace, Name: ingress.Name}, found)
		if err != nil {
			if errors.IsNotFound(err) {
				if err := controllerutil.SetControllerReference(instance, ingress, r.Scheme); err != nil {
					return reconcile.Result{}, err
				}
				logger.Info("Creating a new Ingress",
					"Ingress.Namespace", ingress.Namespace, "Ingress.Name", ingress.Name,
					"Ingress.Annotations", ingress.Annotations, "Ingress.Spec", ingress.Spec)
				return reconcile.Result{}, r.Client.Create(ctx, ingress)
			}
			// Error reading the object - requeue the request.
			return reconcile.Result{}, err
		}

		if reflect.DeepEqual(ingress.Annotations, found.Annotations) &&
			reflect.DeepEqual(ingress.Spec, found.Spec) {

			return reconcile.Result{}, nil
		}

		found.Annotations = ingress.Annotations
		found.Spec = ingress.Spec

		logger.Info("Updating an Ingress",
			"Ingress.Namespace", found.Namespace, "Ingress.Name", found.Name,
			"Ingress.Annotations", found.Annotations, "Ingress.Spec", found.Spec)
		return reconcile.Result{}, r.Client.Update(ctx, found)
	}

	for _, igr := range instance.Status.LoadBalancer.Ingress {
		if igr.IP != "" {
			update := false
			for i, rule := range instance.Spec.Rules {
				if host, found := replaceIp(rule.Host, igr.IP); found {
					instance.Spec.Rules[i].Host = host
					update = true
				}
			}
			for i, tls := range instance.Spec.TLS {
				for j, host := range tls.Hosts {
					if host, found := replaceIp(host, igr.IP); found {
						instance.Spec.TLS[i].Hosts[j] = host
						update = true
					}
				}
			}
			if update {
				logger.Info("Updating an Ingress",
					"Ingress.Namespace", instance.Namespace, "Ingress.Name", instance.Name,
					"Ingress.Spec", instance.Spec)
				return reconcile.Result{}, r.Client.Update(ctx, instance)
			}
			break
		}
	}

	return ctrl.Result{}, nil
}

func (r *IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&extensionsv1beta1.Ingress{}).
		Complete(r)
}

func newIngress(src *extensionsv1beta1.Ingress) *extensionsv1beta1.Ingress {
	annotations := make(map[string]string)
	if src.Annotations != nil {
		for key, val := range src.Annotations {
			switch key {
			case ingressClassAnnotation:
				annotations[ingressClassAnnotation] = strings.TrimSuffix(val, ingressClassSuffix)
			default:
				annotations[key] = val
			}
		}
	}
	ingress := &extensionsv1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:        src.Name + ingressClassSuffix,
			Namespace:   src.Namespace,
			Annotations: annotations,
		},
	}
	src.Spec.DeepCopyInto(&ingress.Spec)

	return ingress
}

func replaceIp(host, ip string) (string, bool) {
	if strings.Contains(host, hostIpPattern) {
		return strings.ReplaceAll(host, hostIpPattern, fmt.Sprintf(".%s.", ip)), true
	}
	return host, false
}
