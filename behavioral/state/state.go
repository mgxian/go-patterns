package state

type state interface {
	handle(ctx *context, v int) string
}

type concreteStateA struct {
}

func (s concreteStateA) handle(ctx *context, v int) string {
	if v < 0 {
		ctx.setState(concreteStateB{})
	}
	return "stateA"
}

type concreteStateB struct {
}

func (s concreteStateB) handle(ctx *context, v int) string {
	if v > 0 {
		ctx.setState(concreteStateA{})
	}
	return "stateB"
}

type context struct {
	state state
	value int
}

func newContext() *context {
	return &context{
		state: concreteStateA{},
	}
}

func (ctx *context) setState(s state) {
	ctx.state = s
}

func (ctx *context) request(v int) string {
	return ctx.state.handle(ctx, v)
}
