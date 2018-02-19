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

	expectedBindings := []Binding{Binding{UserConfig: userConfig1, Service: service1, Ingresses: []Ingress{ingress1}}}

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

func TestGetConsulDto(t *testing.T) {
	clusterIP := "10.11.20.4"
	externalIP := "31.154.22.10"
	service1 := Service{
		Metadata: Metadata{Name: "gateway"},
		Spec:     serviceSpec{ClusterIP: clusterIP, ExternalIP: externalIP},
	}

	annotation := make(map[string]string)
	annotation["kubernetes.io/ingress.class"] = "nginx"

	ingress1 := Ingress{
		Metadata: Metadata{Annotations: annotation},
		Spec: ingressSpec{
			Rules: []ingressRule{
				ingressRule{Host: "payment-service.gateway.service.consul"},
				ingressRule{Host: "availability-service.gateway.service.consul"},
				ingressRule{Host: "invalid-service.gateway.service.com"},
			},
		},
	}
	ingress2 := Ingress{
		Metadata: Metadata{Annotations: annotation},
		Spec: ingressSpec{
			Rules: []ingressRule{
				ingressRule{Host: "login-service.gateway.service.consul"},
				ingressRule{Host: "rating-service.gateway.service.consul"},
			},
		},
	}
	userConfig1 := UserConfig{Name: "config-a", IPType: "clusterIP", ControllerService: "gateway", Annotation: annotation}

	binding := Binding{UserConfig: userConfig1, Service: service1, Ingresses: []Ingress{ingress1, ingress2}}

	actual := binding.GetConsulDto()
	expected := ConsulDto{
		ID:      "gateway",
		Name:    "gateway",
		Tags:    []string{"payment-service", "availability-service", "login-service", "rating-service"},
		Address: clusterIP,
	}
	assert.Equal(t, expected, actual)
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

func TestIsValidHost(t *testing.T) {
	baseDomain := ".base.domain"

	validHost := isValidHost("valid.host.base.domain", baseDomain)
	assert.Equal(t, true, validHost)

	differentDomainHost := isValidHost("invalid.host.different.domain", baseDomain)
	assert.Equal(t, false, differentDomainHost)

	invalidDns := isValidHost("invalid_Dns.host.base.domain", baseDomain)
	assert.Equal(t, false, invalidDns)
}

func TestGetTag(t *testing.T) {
	baseDomain := ".base.domain"

	tag := getTag("valid.host.base.domain", baseDomain)

	assert.Equal(t, "valid.host", tag)
}
