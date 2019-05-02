package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var subject observable
	subject = newSubject()
	subject.attach(concreteObserverA{})
	subject.attach(concreteObserverB{})
	subject.attach(concreteObserverC{})

	assert.Equal(t, subject.notify(), "concreteObserverA concreteObserverB concreteObserverC")
}
