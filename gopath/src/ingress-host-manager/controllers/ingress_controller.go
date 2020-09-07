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
	"regexp"

	"github.com/go-logr/logr"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	hostIpFormat  = ".%s."
	hostIpPattern = regexp.MustCompile(
		`\.(([0–9]|[1–9][0–9]|1[0–9]{2}|2[0–4][0–9]|25[0–5])\.){3}([0–9]|[1–9][0–9]|1[0–9]{2}|2[0–4][0–9]|25[0–5])\.`)
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

	for _, igr := range instance.Status.LoadBalancer.Ingress {
		if igr.IP != "" {
			update := false
			for i, rule := range instance.Spec.Rules {
				if host, changed := replaceIp(rule.Host, igr.IP); changed {
					instance.Spec.Rules[i].Host = host
					update = true
				}
			}
			for i, tls := range instance.Spec.TLS {
				for j, host := range tls.Hosts {
					if host, changed := replaceIp(host, igr.IP); changed {
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

func replaceIp(host, ip string) (string, bool) {
	replace := hostIpPattern.ReplaceAllLiteralString(host, fmt.Sprintf(hostIpFormat, ip))
	return replace, host != replace
}
