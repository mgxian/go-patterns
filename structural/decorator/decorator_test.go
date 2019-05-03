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
	assert.Equal(t, "concreteComponent", aConcreteComponent.operation())

	aConcreteDecoratorA = newConcreteDecoratorA(aConcreteComponent)
	assert.Equal(t, "addedBehavior concreteComponent", aConcreteDecoratorA.operation())

	aConcreteDecoratorB = newConcreteDecoratorB(aConcreteDecoratorA)
	assert.Equal(t, "addedState addedBehavior concreteComponent", aConcreteDecoratorB.operation())
}
