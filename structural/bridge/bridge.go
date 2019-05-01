package bridge

type implementor interface {
	operationImpl() string
}

type concreteImplementorA struct {
}

func (imp concreteImplementorA) operationImpl() string {
	return "concreteImplementorA"
}

type concreteImplementorB struct {
}

func (imp concreteImplementorB) operationImpl() string {
	return "concreteImplementorB"
}

type abstraction struct {
	imp implementor
}

func (a abstraction) operation() string {
	return a.imp.operationImpl()
}

type refinedAbstractionA struct {
	*abstraction
}

func newRefinedAbstractionA(imp implementor) refinedAbstractionA {
	return refinedAbstractionA{
		&abstraction{imp: imp},
	}
}

func (r refinedAbstractionA) operation() string {
	return "refinedAbstractionA " + r.abstraction.operation()
}

type refinedAbstractionB struct {
	*abstraction
}

func newRefinedAbstractionB(imp implementor) refinedAbstractionB {
	return refinedAbstractionB{
		&abstraction{imp: imp},
	}
}

func (r refinedAbstractionB) operation() string {
	return "refinedAbstractionB " + r.abstraction.operation()
}
