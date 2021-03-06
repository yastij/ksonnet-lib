package ksonnet_test

import (
	"testing"

	"github.com/ksonnet/ksonnet-lib/ksonnet-gen/ksonnet"
	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	props := make(map[string]ksonnet.Field)
	props["foo"] = ksonnet.NewLiteralField("name", "integer", "desc", "ref")

	ty := ksonnet.NewType("id", "desc", "group", "ver", "kind", props)

	assert.Equal(t, "id", ty.Identifier())
	assert.Equal(t, "desc", ty.Description())
	assert.Equal(t, "group", ty.Group())
	assert.Equal(t, "group", ty.QualifiedGroup())
	assert.Equal(t, "ver", ty.Version())
	assert.Equal(t, "kind", ty.Kind())
	assert.False(t, ty.IsResource())

	assert.Len(t, ty.Properties(), 1)
}

func TestType_no_group(t *testing.T) {
	props := make(map[string]ksonnet.Field)
	props["foo"] = ksonnet.NewLiteralField("name", "integer", "desc", "ref")

	ty := ksonnet.NewType("id", "desc", "", "ver", "kind", props)

	assert.Equal(t, "core", ty.Group())
	assert.Equal(t, "core", ty.QualifiedGroup())
}
