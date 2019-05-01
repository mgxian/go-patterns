package composite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposite(t *testing.T) {
	var root component
	var aComposite component
	aComposite = newComposite()
	aComposite.add(newLeaf("D"))
	aComposite.add(newLeaf("E"))
	aComposite.add(newLeaf("F"))
	root = newComposite()
	root.add(newLeaf("B"))
	root.add(aComposite)
	root.add(newLeaf("C"))

	assert.Equal(t, root.operation(), "BDEFC")
}
