package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObserver(t *testing.T) {
	var subject observable
	subject = newSubject()
	subject.attach(concreteObserverA{})
	subject.attach(concreteObserverB{})
	subject.attach(concreteObserverC{})

	assert.Equal(t, "concreteObserverA concreteObserverB concreteObserverC", subject.notify())
}
