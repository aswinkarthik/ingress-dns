package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIPAddressReturnsClusterIP(t *testing.T) {
	binding := &Binding{
		IngressConfig: IngressConfig{
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
		IngressConfig: IngressConfig{
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
		IngressConfig: IngressConfig{
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
		IngressConfig: IngressConfig{
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
