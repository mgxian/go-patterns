package template_method

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateMethod(t *testing.T) {
	cc := newConcreteClass()
	ac := newAbstractClass(cc)
	assert.Equal(t, "operationA operationB operationC", ac.templateMethod())
}
