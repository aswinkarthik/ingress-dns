package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBindings(t *testing.T) {
	binding1 := Binding{
		UserConfig: UserConfig{
			Name:   "nginx-internal",
			IPType: "internalIP",
		},
		Ingress: Ingress{
			Spec: ingressSpec{
				Rules: []ingressRule{
					ingressRule{Host: "abc.internal"},
				},
			},
		},
		Service: Service{
			Spec: serviceSpec{
				ExternalIP: "10.2.3.4",
			},
		},
	}
	binding2 := Binding{
		UserConfig: UserConfig{
			Name:   "nginx-internal",
			IPType: "internalIP",
		},
		Ingress: Ingress{
			Spec: ingressSpec{Rules: []ingressRule{
				ingressRule{Host: "def.internal"},
			},
			},
		},
		Service: Service{
			Spec: serviceSpec{
				ExternalIP: "10.2.3.4",
			},
		},
	}
	binding3 := Binding{
		UserConfig: UserConfig{
			Name:   "nginx-external",
			IPType: "externalIP",
		},
		Ingress: Ingress{
			Spec: ingressSpec{Rules: []ingressRule{
				ingressRule{Host: "ijk.com"},
			},
			},
		},
		Service: Service{
			Spec: serviceSpec{
				ExternalIP: "1.2.3.4",
			},
		},
	}

	groupedBindings := groupBindings([]Binding{binding1, binding2, binding3})
	expectedGroupedBindings := make(map[string][]Binding, 2)
	expectedGroupedBindings["nginx-internal"] = []Binding{binding1, binding2}
	expectedGroupedBindings["nginx-external"] = []Binding{binding3}

	assert.Equal(t, expectedGroupedBindings, groupedBindings)
}

func TestGetTags(t *testing.T) {
	hosts := []string{"service-a.service.consul", "service-b.service.consul", "service-c.skipped.domain"}
	domain := ".service.consul"
	expectedTags := []string{"service-a", "service-b"}

	assert.Equal(t, expectedTags, getTags(hosts, domain))
}

func TestGetTagsShouldSkipInvalidHosts(t *testing.T) {
	hosts := []string{"wrong_host.service.consul"}
	domain := ".service.consul"
	expectedTags := []string{}

	assert.Equal(t, expectedTags, getTags(hosts, domain))
}
