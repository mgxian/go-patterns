package chain_of_responsibility

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
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

	assert.Equal(t, "handlerA 111 handle", handlerA1.handleRequest(111))
	assert.Equal(t, "handlerB 113 handle", handlerA1.handleRequest(113))
	assert.Equal(t, "handlerA 110 handle", handlerA1.handleRequest(110))
	assert.Equal(t, "handlerB 112 handle", handlerA1.handleRequest(112))
	assert.Equal(t, "no handler can handle this", handlerA1.handleRequest(114))
}
