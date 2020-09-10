package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

func TestController(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	const (
		name       = "test"
		namespace  = "default"
		hostMask   = "test.0.0.0.0.xip.io"
		hostWithIP = "test.1.2.3.4.xip.io"
		ip         = "1.2.3.4"
	)

	igr := &extensionsv1beta1.Ingress{
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
		Status: extensionsv1beta1.IngressStatus{
			LoadBalancer: corev1.LoadBalancerStatus{
				Ingress: []corev1.LoadBalancerIngress{
					{
						IP: ip,
					},
				},
			},
		},
	}

	// Objects to track in the fake client.
	objs := []runtime.Object{
		igr,
	}
	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(extensionsv1beta1.SchemeGroupVersion, igr)
	// Create a fake client to mock API calls.
	cl := fake.NewFakeClientWithScheme(s, objs...)

	// Create a ReconcileOpenApiService object with the scheme and fake client.
	r := &IngressReconciler{Client: cl, Log: logf.Log, Scheme: s}

	// Mock request to simulate Reconcile() being called on an event for a
	// watched resource .
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      name,
			Namespace: namespace,
		},
	}
	_, err := r.Reconcile(req)

	assert.Nil(t, err)

	igr = &extensionsv1beta1.Ingress{}
	err = cl.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, igr)

	assert.Nil(t, err)
	assert.Equal(t, hostWithIP, igr.Spec.Rules[0].Host)
	assert.Equal(t, hostWithIP, igr.Spec.TLS[0].Hosts[0])
}

func TestController_notFound(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	var (
		name      = "test"
		namespace = "testns"
	)

	// Objects to track in the fake client.
	objs := []runtime.Object{}
	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(extensionsv1beta1.SchemeGroupVersion, &extensionsv1beta1.Ingress{})
	// Create a fake client to mock API calls.
	cl := fake.NewFakeClientWithScheme(s, objs...)

	// Create a ReconcileOpenApiService object with the scheme and fake client.
	r := &IngressReconciler{Client: cl, Log: logf.Log, Scheme: s}

	// Mock request to simulate Reconcile() being called on an event for a
	// watched resource .
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      name,
			Namespace: namespace,
		},
	}
	_, err := r.Reconcile(req)

	assert.Nil(t, err)

	igr := &extensionsv1beta1.Ingress{}
	err = cl.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, igr)

	assert.Error(t, err)
	assert.True(t, errors.IsNotFound(err))
}
