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
	assert.Equal(t, "ProductA1", productA.show())
	assert.Equal(t, "ProductB1", productB.show())
}

func TestCreateProductByFactory2(t *testing.T) {
	var factory AbstractFactory
	var productA ProductA
	var productB ProductB
	factory = ConcreteFactory2{}
	productA = factory.CreateProductA()
	productB = factory.CreateProductB()
	assert.Equal(t, "ProductA2", productA.show())
	assert.Equal(t, "ProductB2", productB.show())
}
