package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator(t *testing.T) {
	var aIterator iterator

	aAggregate := newConcreteAggregate()
	aIterator = aAggregate.createIterator()
	aAggregate.add(1)
	aAggregate.add(2)
	aAggregate.add(3)
	aAggregate.add("a")

	assert.Equal(t, 1, aIterator.currentItem())
	assert.Equal(t, true, aIterator.hasNext())
	aIterator.next()

	assert.Equal(t, 2, aIterator.currentItem())
	assert.Equal(t, true, aIterator.hasNext())
	aIterator.next()

	assert.Equal(t, 3, aIterator.currentItem())
	assert.Equal(t, true, aIterator.hasNext())
	aIterator.next()

	assert.Equal(t, "a", aIterator.currentItem())
	assert.Equal(t, false, aIterator.hasNext())
}
