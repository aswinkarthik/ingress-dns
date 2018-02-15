package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBindings(t *testing.T) {
	service1 := Service{Metadata: Metadata{Name: "controller-service-a"}}
	services := []Service{service1}
	serviceList := ServiceList{Items: services}

	controllerAnnotation := make(map[string]string)
	controllerAnnotation["my-annotation"] = "value"

	differentAnnotation := make(map[string]string)
	differentAnnotation["my-annotation"] = "different"

	ingress1 := Ingress{Metadata: Metadata{Annotations: controllerAnnotation}}
	ingress2 := Ingress{Metadata: Metadata{Annotations: differentAnnotation}}
	ingresses := []Ingress{ingress1, ingress2}
	ingressList := IngressList{Items: ingresses}

	userConfig1 := UserConfig{Name: "config-a", IPType: "clusterIP", ControllerService: "controller-service-a", Annotation: controllerAnnotation}
	userConfigs := []UserConfig{userConfig1}

	expectedBindings := []BindingV2{BindingV2{UserConfig: userConfig1, Service: service1, Ingresses: []Ingress{ingress1}}}

	assert.Equal(t, expectedBindings, NewBindings(serviceList, ingressList, userConfigs))
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
