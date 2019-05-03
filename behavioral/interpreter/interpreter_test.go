package interpreter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterpreter(t *testing.T) {
	aExpression := "a b c - +"
	variables := make(map[string]expression)
	variables["a"] = &number{10}
	variables["b"] = &number{20}
	variables["c"] = &number{8}
	aEvaluator := newEvaluator(aExpression, variables)
	actual := aEvaluator.execute()
	expeted := 22
	assert.Equal(t, expeted, actual)
}
