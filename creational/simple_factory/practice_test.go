package simple_factory

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestCreateCircle(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	s := newShape(Circle)
	s.draw()
	s.erase()

	assert.Equal(t, "Draw circle\nErase circle\n", outputWriter.(*strings.Builder).String())
}

func TestCreateSquare(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	s := newShape(Square)
	s.draw()
	s.erase()

	assert.Equal(t, "Draw square\nErase square\n", outputWriter.(*strings.Builder).String())
}

func TestCreateTriangle(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	s := newShape(Triangle)
	s.draw()
	s.erase()

	assert.Equal(t, "Draw triangle\nErase triangle\n", outputWriter.(*strings.Builder).String())
}
