package strategy

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlane(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var aPlane plane = &helicopter{}
	aPlane.takeOff()
	aPlane.fly()

	expected := "helicopter take off\nhelicopter fly\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	aPlane = &airplane{}
	aPlane.takeOff()
	aPlane.fly()

	expected = "airplane take off\nairplane fly\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	aPlane = &harrier{}
	aPlane.takeOff()
	aPlane.fly()

	expected = "harrier take off\nharrier fly\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
}
