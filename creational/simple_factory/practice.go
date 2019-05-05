package simple_factory

import (
	"fmt"
	"io"
	"os"
)

// collect stdout
var outputWriter io.Writer = os.Stdout

type shape interface {
	draw()
	erase()
}

type circle struct {
}

func newCircle() *circle {
	return &circle{}
}

func (c *circle) draw() {
	fmt.Fprintln(outputWriter, "Draw circle")
}

func (c *circle) erase() {
	fmt.Fprintln(outputWriter, "Erase circle")
}

type square struct {
}

func newSquare() *square {
	return &square{}
}

func (s *square) draw() {
	fmt.Fprintln(outputWriter, "Draw square")
}

func (s *square) erase() {
	fmt.Fprintln(outputWriter, "Erase square")
}

type triangle struct {
}

func newTriangle() *triangle {
	return &triangle{}
}

func (t *triangle) draw() {
	fmt.Fprintln(outputWriter, "Draw triangle")
}

func (t *triangle) erase() {
	fmt.Fprintln(outputWriter, "Erase triangle")
}

type shapeType int

// shpae type enum
const (
	Circle shapeType = iota
	Square
	Triangle
)

func newShape(t shapeType) shape {
	switch t {
	case Circle:
		return newCircle()
	case Square:
		return newSquare()
	case Triangle:
		return newTriangle()
	default:
		return nil
	}
}
