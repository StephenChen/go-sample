package base

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	assert1 := assert.New(t)
	var object interface{}

	// assert equality
	assert1.Equal(123, 123, "they should be equal")

	// assert inequality
	assert1.NotEqual(123, 456, "they should not be equal")

	// assert fot nil (good for errors)
	assert1.Nil(object)

	// assert for not nil (good when you expect something)
	if assert1.NotNil(object) {
		assert1.Equal("Something", object)
	}
}
