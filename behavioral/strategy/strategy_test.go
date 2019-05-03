package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrategy(t *testing.T) {
	ctx := newContext(concreteStrategyA{})
	assert.Equal(t, "concreteStrategyA", ctx.algorithm())

	ctx.setStrategy(concreteStrategyB{})
	assert.Equal(t, "concreteStrategyB", ctx.algorithm())
}
