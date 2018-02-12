package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServiceMap(t *testing.T) {
	serviceList := &ServiceList{
		Items: []Service{
			Service{Metadata: Metadata{Name: "service-a"}},
			Service{Metadata: Metadata{Name: "service-b"}},
		},
	}
	actualServiceMap := serviceList.GetServiceMap()

	expectedServiceMap := make(map[string]Service, 2)
	expectedServiceMap["service-a"] = Service{
		Metadata: Metadata{Name: "service-a"},
	}
	expectedServiceMap["service-b"] = Service{
		Metadata: Metadata{Name: "service-b"},
	}

	assert.Equal(t, actualServiceMap, expectedServiceMap)
}
