package memento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemento(t *testing.T) {
	aOriginator := newOriginator("A")
	assert.Equal(t, "A", aOriginator.getState())

	aCaretaker := newCaretaker(aOriginator.createMemento())
	aOriginator.setState("B")
	aCaretaker.setMemento(aOriginator.createMemento())

	aOriginator.setState("C")
	assert.Equal(t, "C", aOriginator.getState())

	aOriginator.restoreMemento(aCaretaker.getMemento())
	assert.Equal(t, "B", aOriginator.getState())
}
