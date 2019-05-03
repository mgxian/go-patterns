package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacade(t *testing.T) {
	aSubSystemA := subSystemA{}
	aSubSystemB := subSystemB{}
	aSubSystemC := subSystemC{}
	aFacade := newFacade(aSubSystemA, aSubSystemB, aSubSystemC)
	assert.Equal(t, "ABC", aFacade.operation())
}
