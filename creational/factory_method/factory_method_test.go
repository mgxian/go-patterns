package factory_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProductA(t *testing.T) {
	var factory Factory
	var product Product
	factory = ConcreteAFactory{}
	product = factory.CreateProduct()
	assert.Equal(t, product.show(), "ProductA")
}

func TestCreateProductB(t *testing.T) {
	var factory Factory
	var product Product
	factory = ConcreteBFactory{}
	product = factory.CreateProduct()
	assert.Equal(t, product.show(), "ProductB")
}
