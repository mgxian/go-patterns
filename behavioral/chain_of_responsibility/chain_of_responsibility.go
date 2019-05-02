package chain_of_responsibility

import "fmt"

type handler interface {
	handleRequest(int) string
	setSuccessor(handler)
}

type abstractHandler struct {
	successor handler
}

func newAbstractHandler() *abstractHandler {
	return &abstractHandler{}
}

func (ah *abstractHandler) setSuccessor(aSuccessor handler) {
	ah.successor = aSuccessor
}

type concreteHandlerA struct {
	abstractHandler
	hid int
}

func newConcreteHandlerA(hid int) *concreteHandlerA {
	return &concreteHandlerA{
		hid: hid,
	}
}

func (h *concreteHandlerA) handleRequest(n int) string {
	if h.hid == n {
		return fmt.Sprintf("handlerA %d handle", h.hid)
	}

	if h.successor == nil {
		return fmt.Sprintf("no handler can handle this")
	}

	return h.successor.handleRequest(n)
}

type concreteHandlerB struct {
	abstractHandler
	hid int
}

func newConcreteHandlerB(hid int) *concreteHandlerB {
	return &concreteHandlerB{
		hid: hid,
	}
}

func (h *concreteHandlerB) handleRequest(n int) string {
	if h.hid == n {
		return fmt.Sprintf("handlerB %d handle", h.hid)
	}

	if h.successor == nil {
		return fmt.Sprintf("no handler can handle this")
	}

	return h.successor.handleRequest(n)
}
