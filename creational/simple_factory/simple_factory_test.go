package simple_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProductA(t *testing.T) {
	factory := Factory{}
	product := factory.GetProduct("A")
	assert.Equal(t, product.show(), "ProductA")
}

func TestCreateProductB(t *testing.T) {
	factory := Factory{}
	product := factory.GetProduct("B")
	assert.Equal(t, product.show(), "ProductB")
}

func TestCreateUnknownProduct(t *testing.T) {
	factory := Factory{}
	product := factory.GetProduct("Unkown")
	assert.Nil(t, product)
}