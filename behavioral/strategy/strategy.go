package strategy

type strategy interface {
	algorithm() string
}

type concreteStrategyA struct {
}

func (s concreteStrategyA) algorithm() string {
	return "concreteStrategyA"
}

type concreteStrategyB struct {
}

func (s concreteStrategyB) algorithm() string {
	return "concreteStrategyB"
}

type context struct {
	aStrategy strategy
}

func newContext(s strategy) *context {
	return &context{
		aStrategy: s,
	}
}

func (ctx *context) setStrategy(s strategy) {
	ctx.aStrategy = s
}

func (ctx context) algorithm() string {
	return ctx.aStrategy.algorithm()
}
