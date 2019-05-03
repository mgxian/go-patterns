package simple_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProductA(t *testing.T) {
	var product Product
	factory := Factory{}
	product = factory.CreateProduct(ProductA)
	assert.Equal(t, "ProductA", product.show())
}

func TestCreateProductB(t *testing.T) {
	var product Product
	factory := Factory{}
	product = factory.CreateProduct(ProductB)
	assert.Equal(t, "ProductB", product.show())
}

func TestCreateUnknownProduct(t *testing.T) {
	var product Product
	factory := Factory{}
	product = factory.CreateProduct(ProductType(3))
	assert.Nil(t, product)
}