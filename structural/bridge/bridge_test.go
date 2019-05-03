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
	assert.Equal(t, "refinedAbstractionA concreteImplementorA", result)
}

func TestRefinedAbstractionAWithConcreteImplementorB(t *testing.T) {
	var imp implementor
	imp = concreteImplementorB{}
	ra := newRefinedAbstractionA(imp)
	result := ra.operation()
	assert.Equal(t, "refinedAbstractionA concreteImplementorB", result)
}

func TestRefinedAbstractionBWithConcreteImplementorA(t *testing.T) {
	var imp implementor
	imp = concreteImplementorA{}
	ra := newRefinedAbstractionB(imp)
	result := ra.operation()
	assert.Equal(t, "refinedAbstractionB concreteImplementorA", result)
}
