package prototype

import (
	"testing"

	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	var prototype Prototype
	prototype = ConcretePrototype{}
	newPrototype := prototype.Clone()
	assert.NotEqual(t, unsafe.Pointer(&newPrototype), unsafe.Pointer(&prototype))
}
