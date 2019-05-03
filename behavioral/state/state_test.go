package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestState(t *testing.T) {
	ctx := newContext()
	assert.Equal(t, "stateA", ctx.request(-1))
	assert.Equal(t, "stateB", ctx.request(1))
}
