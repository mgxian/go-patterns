package chain_of_responsibility

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var handlerA1 handler
	var handlerA2 handler
	var handlerB1 handler
	var handlerB2 handler

	handlerA1 = newConcreteHandlerA(110)
	handlerA2 = newConcreteHandlerA(111)
	handlerB1 = newConcreteHandlerB(112)
	handlerB2 = newConcreteHandlerB(113)

	handlerA1.setSuccessor(handlerA2)
	handlerA2.setSuccessor(handlerB1)
	handlerB1.setSuccessor(handlerB2)

	assert.Equal(t, handlerA1.handleRequest(111), "handlerA 111 handle")
	assert.Equal(t, handlerA1.handleRequest(113), "handlerB 113 handle")
	assert.Equal(t, handlerA1.handleRequest(110), "handlerA 110 handle")
	assert.Equal(t, handlerA1.handleRequest(112), "handlerB 112 handle")
	assert.Equal(t, handlerA1.handleRequest(114), "no handler can handle this")
}
