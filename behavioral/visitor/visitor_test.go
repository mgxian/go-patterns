package visitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisitor(t *testing.T) {
	var v visitor
	el := newElementList()
	v = concreteVisitorA{}

	el.add(newConcreteElementA("go"))
	el.add(newConcreteElementA("rust"))
	el.add(newConcreteElementA("c"))

	el.add(newConcreteElementB("python"))
	el.add(newConcreteElementB("js"))
	el.add(newConcreteElementB("ruby"))

	expected := []string{
		"concreteVisitorA concreteElementA go operationA",
		"concreteVisitorA concreteElementA rust operationA",
		"concreteVisitorA concreteElementA c operationA",
		"concreteVisitorA concreteElementB python operationB",
		"concreteVisitorA concreteElementB js operationB",
		"concreteVisitorA concreteElementB ruby operationB",
	}
	assert.Equal(t, el.accept(v), expected)

	v = concreteVisitorB{}
	expected = []string{
		"concreteVisitorB concreteElementA go operationA",
		"concreteVisitorB concreteElementA rust operationA",
		"concreteVisitorB concreteElementA c operationA",
		"concreteVisitorB concreteElementB python operationB",
		"concreteVisitorB concreteElementB js operationB",
		"concreteVisitorB concreteElementB ruby operationB",
	}
	assert.Equal(t, el.accept(v), expected)
}
