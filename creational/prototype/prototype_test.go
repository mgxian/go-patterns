package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"unsafe"
)

func TestPrototype(t *testing.T) {
	var prototype Prototype
	prototype = ConcretePrototypeA{name: "A"}
	newPrototype := prototype.Clone()	
	assert.NotEqual(t, unsafe.Pointer(&newPrototype), unsafe.Pointer(&prototype))
}
