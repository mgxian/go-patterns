package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ctx := newContext(concreteStrategyA{})
	assert.Equal(t, ctx.algorithm(), "concreteStrategyA")

	ctx.setStrategy(concreteStrategyB{})
	assert.Equal(t, ctx.algorithm(), "concreteStrategyB")
}
