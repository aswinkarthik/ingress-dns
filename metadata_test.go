package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsAnnotations(t *testing.T) {
	annotations := make(map[string]string, 2)
	annotations["annotation-1"] = "value-1"
	annotations["annotation-2"] = "value-2"

	metadata := &Metadata{
		Annotations: annotations,
	}

	present := make(map[string]string, 1)
	present["annotation-2"] = "value-2"

	notPresent := make(map[string]string, 1)
	notPresent["annotation-3"] = "missing"

	assert.Equal(t, metadata.ContainsAnnotations(present), true)
	assert.Equal(t, metadata.ContainsAnnotations(notPresent), false)
}
