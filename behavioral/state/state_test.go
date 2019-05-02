package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ctx := newContext()
	assert.Equal(t, ctx.request(-1), "stateA")
	assert.Equal(t, ctx.request(1), "stateB")
}
