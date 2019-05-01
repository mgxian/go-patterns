package simple_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProductA(t *testing.T) {
	var product Product
	factory := Factory{}
	product = factory.CreateProduct(ProductA)
	assert.Equal(t, product.show(), "ProductA")
}

func TestCreateProductB(t *testing.T) {
	var product Product
	factory := Factory{}
	product = factory.CreateProduct(ProductB)
	assert.Equal(t, product.show(), "ProductB")
}

func TestCreateUnknownProduct(t *testing.T) {
	var product Product
	factory := Factory{}
	product = factory.CreateProduct(ProductType(3))
	assert.Nil(t, product)
}