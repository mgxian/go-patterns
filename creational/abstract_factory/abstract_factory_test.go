package abstract_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProductByFactory1(t *testing.T) {
	var factory AbstractFactory
	var productA ProductA
	var productB ProductB
	factory = ConcreteFactory1{}
	productA = factory.CreateProductA()
	productB = factory.CreateProductB()
	assert.Equal(t, productA.show(), "ProductA1")
	assert.Equal(t, productB.show(), "ProductB1")
}

func TestCreateProductByFactory2(t *testing.T) {
	var factory AbstractFactory
	var productA ProductA
	var productB ProductB
	factory = ConcreteFactory2{}
	productA = factory.CreateProductA()
	productB = factory.CreateProductB()
	assert.Equal(t, productA.show(), "ProductA2")
	assert.Equal(t, productB.show(), "ProductB2")
}