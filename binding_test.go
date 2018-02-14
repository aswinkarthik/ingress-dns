package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBindings(t *testing.T) {
	services := []Service{
		Service{Metadata: Metadata{Name: "constroller-service-a"}},
	}

	controllerAnnotation := make(map[string]string)

	ingresses := []Ingress{
		Ingress{Metadata: Metadata{Annotations: controllerAnnotation}},
	}

	userConfigs := []UserConfig{
		UserConfig{Name: "config-a", IPType: "clusterIP", ControllerService: "controller-service-a", Annotation: controllerAnnotation},
	}

	expectedBindings := []Binding{}

	assert.Equal(t, expectedBindings, NewBindings(services, ingresses, userConfigs))
}

func TestGetIPAddressReturnsClusterIP(t *testing.T) {
	binding := &Binding{
		UserConfig: UserConfig{
			IPType: "clusterIP",
		},
		Service: Service{
			Spec: serviceSpec{
				ClusterIP:  "10.7.0.1",
				ExternalIP: "31.24.56.112",
			},
		},
	}
	assert.Equal(t, binding.GetIPAddress(), "10.7.0.1")
}

func TestGetIPAddressReturnsExternalIP(t *testing.T) {
	binding := &Binding{
		UserConfig: UserConfig{
			IPType: "externalIP",
		},
		Service: Service{
			Spec: serviceSpec{
				ClusterIP:  "10.7.0.1",
				ExternalIP: "31.24.56.112",
			},
		},
	}
	assert.Equal(t, binding.GetIPAddress(), "31.24.56.112")
}

func TestGetIPAddressReturnsEmptyString(t *testing.T) {
	binding := &Binding{
		UserConfig: UserConfig{
			IPType: "loadBalancerIP",
		},
		Service: Service{
			Spec: serviceSpec{
				ClusterIP:  "10.7.0.1",
				ExternalIP: "31.24.56.112",
			},
		},
	}
	assert.Equal(t, binding.GetIPAddress(), "")
}

func TestGetIdAndName(t *testing.T) {
	binding := &Binding{
		UserConfig: UserConfig{
			Name: "user-defined-name",
		},
	}

	assert.Equal(t, binding.GetId(), "user-defined-name")
	assert.Equal(t, binding.GetName(), "user-defined-name")
}

func TestGetHosts(t *testing.T) {
	binding := &Binding{
		Ingress: Ingress{
			Spec: ingressSpec{
				Rules: []ingressRule{
					ingressRule{Host: "abc.com"},
				},
			},
		},
	}

	assert.Equal(t, binding.GetHosts(), []string{"abc.com"})
}
