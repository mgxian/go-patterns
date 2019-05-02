package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var cmd command
	var aInvoker *invoker

	cmd = newConcreteCommandA()
	aInvoker = newInvoker(cmd)
	assert.Equal(t, aInvoker.call(), "A")

	cmd = newConcreteCommandB()
	aInvoker.setCommand(cmd)
	assert.Equal(t, aInvoker.call(), "B")
}
