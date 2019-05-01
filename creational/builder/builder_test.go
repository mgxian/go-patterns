package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildProduct(t *testing.T) {
	var product Product
	var aBuilder builder
	var aDirector *director
	aBuilder = newConcreteBuilder()
	aDirector = newDirector(aBuilder)
	product = aDirector.construct()
	assert.Equal(t, product.getResult(), "ABC")
}
