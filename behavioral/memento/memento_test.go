package memento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	aOriginator := newOriginator("A")
	assert.Equal(t, aOriginator.getState(), "A")

	aCaretaker := newCaretaker(aOriginator.createMemento())
	aOriginator.setState("B")
	aCaretaker.setMemento(aOriginator.createMemento())

	aOriginator.setState("C")
	assert.Equal(t, aOriginator.getState(), "C")

	aOriginator.restoreMemento(aCaretaker.getMemento())
	assert.Equal(t, aOriginator.getState(), "B")
}
