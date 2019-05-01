package simple_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProductA(t *testing.T) {
	var product Product
	factory := Factory{}
	product = factory.CreateProduct("A")
	assert.Equal(t, product.show(), "ProductA")
}

func TestCreateProductB(t *testing.T) {
	var product Product
	factory := Factory{}
	product = factory.CreateProduct("B")
	assert.Equal(t, product.show(), "ProductB")
}

func TestCreateUnknownProduct(t *testing.T) {
	var product Product
	factory := Factory{}
	product = factory.CreateProduct("Unkown")
	assert.Nil(t, product)
}