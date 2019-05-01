package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefinedAbstractionAWithConcreteImplementorA(t *testing.T) {
	var imp implementor
	imp = concreteImplementorA{}
	ra := newRefinedAbstractionA(imp)
	result := ra.operation()
	assert.Equal(t, result, "refinedAbstractionA concreteImplementorA")
}

func TestRefinedAbstractionAWithConcreteImplementorB(t *testing.T) {
	var imp implementor
	imp = concreteImplementorB{}
	ra := newRefinedAbstractionA(imp)
	result := ra.operation()
	assert.Equal(t, result, "refinedAbstractionA concreteImplementorB")
}

func TestRefinedAbstractionBWithConcreteImplementorA(t *testing.T) {
	var imp implementor
	imp = concreteImplementorA{}
	ra := newRefinedAbstractionB(imp)
	result := ra.operation()
	assert.Equal(t, result, "refinedAbstractionB concreteImplementorA")
}
