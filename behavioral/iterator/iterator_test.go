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

	assert.Equal(t, aIterator.currentItem(), 1)
	assert.Equal(t, aIterator.hasNext(), true)
	aIterator.next()

	assert.Equal(t, aIterator.currentItem(), 2)
	assert.Equal(t, aIterator.hasNext(), true)
	aIterator.next()

	assert.Equal(t, aIterator.currentItem(), 3)
	assert.Equal(t, aIterator.hasNext(), true)
	aIterator.next()

	assert.Equal(t, aIterator.currentItem(), "a")
	assert.Equal(t, aIterator.hasNext(), false)
}
