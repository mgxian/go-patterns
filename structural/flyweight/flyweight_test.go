package flyweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"unsafe"
)

func TestFlyweight(t *testing.T) {
	factory := newFlyweightFactory()
	flyweightA := factory.GetFlyweight("A")
	newFlyweightA := factory.GetFlyweight("A")
	assert.Equal(t, unsafe.Pointer(flyweightA), unsafe.Pointer(newFlyweightA))
}
