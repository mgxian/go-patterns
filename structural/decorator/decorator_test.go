package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	var aConcreteComponent component
	var aConcreteDecoratorA component
	var aConcreteDecoratorB component

	aConcreteComponent = concreteComponent{}
	assert.Equal(t, aConcreteComponent.operation(), "concreteComponent")

	aConcreteDecoratorA = newConcreteDecoratorA(aConcreteComponent)
	assert.Equal(t, aConcreteDecoratorA.operation(), "addedBehavior concreteComponent")

	aConcreteDecoratorB = newConcreteDecoratorB(aConcreteDecoratorA)
	assert.Equal(t, aConcreteDecoratorB.operation(), "addedState addedBehavior concreteComponent")
}
