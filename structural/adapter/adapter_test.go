package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapter(t *testing.T) {
	var aAdaptee adaptee
	var aTarget target

	aAdaptee = newAdaptee()
	aTarget = newAdapter(aAdaptee)
	resp := aTarget.request()
	assert.Equal(t, "adapter response + raw response", resp)
}
