package visitor

type element interface {
	accept(v visitor) string
}

type concreteElementA struct {
	name string
}

func newConcreteElementA(name string) concreteElementA {
	return concreteElementA{
		name: name,
	}
}

func (e concreteElementA) accept(v visitor) string {
	return v.visitA(e)
}

func (e concreteElementA) operationA() string {
	return "concreteElementA " + e.name + " operationA"
}

type concreteElementB struct {
	name string
}

func newConcreteElementB(name string) concreteElementB {
	return concreteElementB{
		name: name,
	}
}

func (e concreteElementB) accept(v visitor) string {
	return v.visitB(e)
}

func (e concreteElementB) operationB() string {
	return "concreteElementB " + e.name + " operationB"
}

type visitor interface {
	visitA(e concreteElementA) string
	visitB(e concreteElementB) string
}

type concreteVisitorA struct {
}

func (v concreteVisitorA) visitA(e concreteElementA) string {
	return "concreteVisitorA " + e.operationA()
}

func (v concreteVisitorA) visitB(e concreteElementB) string {
	return "concreteVisitorA " + e.operationB()
}

type concreteVisitorB struct {
}

func (v concreteVisitorB) visitA(e concreteElementA) string {
	return "concreteVisitorB " + e.operationA()
}

func (v concreteVisitorB) visitB(e concreteElementB) string {
	return "concreteVisitorB " + e.operationB()
}

type elementList struct {
	data []element
}

func newElementList() *elementList {
	return &elementList{
		data: make([]element, 0),
	}
}

func (el *elementList) add(e element) {
	el.data = append(el.data, e)
}

func (el *elementList) accept(v visitor) (result []string) {
	for _, e := range el.data {
		result = append(result, e.accept(v))
	}
	return result
}
