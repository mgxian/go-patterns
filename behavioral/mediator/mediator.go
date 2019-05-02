package mediator

const (
	colleagueA int = iota
	colleagueB
	colleagueC
)

type colleague struct {
	m mediator
}

func (c *colleague) setMediator(m mediator) {
	c.m = m
}

type concreteColleagueA struct {
	colleague
}

func newConcreteColleagueA() *concreteColleagueA {
	return &concreteColleagueA{}
}

func (c concreteColleagueA) methodA() string {
	return "colleagueA_methodA"
}

func (c concreteColleagueA) operation() string {
	return c.m.operation(colleagueA)
}

type concreteColleagueB struct {
	colleague
}

func newConcreteColleagueB() *concreteColleagueB {
	return &concreteColleagueB{}
}

func (c concreteColleagueB) methodB() string {
	return "colleagueB_methodB"
}

func (c concreteColleagueB) operation() string {
	return c.m.operation(colleagueB)
}

type concreteColleagueC struct {
	colleague
}

func newConcreteColleagueC() *concreteColleagueC {
	return &concreteColleagueC{}
}

func (c concreteColleagueC) methodC() string {
	return "colleagueC_methodC"
}

func (c concreteColleagueC) operation() string {
	return c.m.operation(colleagueC)
}

type mediator interface {
	operation(colleagueType int) string
}

type concreteMediator struct {
	colleagueA *concreteColleagueA
	colleagueB *concreteColleagueB
	colleagueC *concreteColleagueC
}

func newConcreteMediator() *concreteMediator {
	return &concreteMediator{}
}

func (m *concreteMediator) operation(colleagueType int) string {
	result := ""
	switch colleagueType {
	case colleagueA:
		result = m.colleagueB.methodB()
	case colleagueB:
		result = m.colleagueA.methodA() + " " + m.colleagueC.methodC()
	case colleagueC:
		result = m.colleagueA.methodA()
	default:
	}

	return result
}
