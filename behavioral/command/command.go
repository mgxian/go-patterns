package command

type command interface {
	execute() string
}

type invoker struct {
	c command
}

func newInvoker(c command) *invoker {
	return &invoker{
		c: c,
	}
}

func (i invoker) call() string {
	return i.c.execute()
}

func (i *invoker) setCommand(c command) {
	i.c = c
}

type receiverA struct {
	name string
}

func newReceiverA(name string) *receiverA {
	return &receiverA{
		name: name,
	}
}

func (r receiverA) action() string {
	return r.name
}

type receiverB struct {
	name string
}

func newReceiverB(name string) *receiverB {
	return &receiverB{
		name: name,
	}
}

func (r receiverB) action() string {
	return r.name
}

type concreteCommandA struct {
	aReceiverA *receiverA
}

func newConcreteCommandA() *concreteCommandA {
	return &concreteCommandA{
		aReceiverA: newReceiverA("A"),
	}
}

func (c concreteCommandA) execute() string {
	return c.aReceiverA.action()
}

type concreteCommandB struct {
	aReceiverB *receiverB
}

func newConcreteCommandB() *concreteCommandB {
	return &concreteCommandB{
		aReceiverB: newReceiverB("B"),
	}
}

func (c concreteCommandB) execute() string {
	return c.aReceiverB.action()
}
