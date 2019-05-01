package decorator

type component interface {
	operation() string
}

type concreteComponent struct {
}

func (c concreteComponent) operation() string {
	return "concreteComponent"
}

type decorator struct {
	c component
}

func (d decorator) operation() string {
	return ""
}

type concreteDecoratorA struct {
	decorator
}

func newConcreteDecoratorA(c component) *concreteDecoratorA {
	return &concreteDecoratorA{
		decorator{c: c},
	}
}

func (cd concreteDecoratorA) operation() string {
	return cd.addedBehavior() + cd.c.operation()
}

func (cd concreteDecoratorA) addedBehavior() string {
	return "addedBehavior "
}

type concreteDecoratorB struct {
	decorator
}

func newConcreteDecoratorB(c component) *concreteDecoratorB {
	return &concreteDecoratorB{
		decorator{c: c},
	}
}

func (cd concreteDecoratorB) operation() string {
	return cd.addedState() + cd.c.operation()
}

func (cd concreteDecoratorB) addedState() string {
	return "addedState "
}
