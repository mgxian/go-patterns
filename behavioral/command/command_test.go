package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommand(t *testing.T) {
	var cmd command
	var aInvoker *invoker

	cmd = newConcreteCommandA()
	aInvoker = newInvoker(cmd)
	assert.Equal(t, "A", aInvoker.call())

	cmd = newConcreteCommandB()
	aInvoker.setCommand(cmd)
	assert.Equal(t, "B", aInvoker.call())
}
