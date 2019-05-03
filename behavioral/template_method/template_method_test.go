package template_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	cc := newConcreteClass()
	ac := newAbstractClass(cc)
	assert.Equal(t, "operationA operationB operationC", ac.templateMethod())
}
