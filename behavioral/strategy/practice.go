package strategy

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type plane interface {
	takeOff()
	fly()
}

type helicopter struct {
}

func (h *helicopter) takeOff() {
	fmt.Fprintln(outputWriter, "helicopter take off")
}

func (h *helicopter) fly() {
	fmt.Fprintln(outputWriter, "helicopter fly")
}

type airplane struct {
}

func (a *airplane) takeOff() {
	fmt.Fprintln(outputWriter, "airplane take off")
}

func (a *airplane) fly() {
	fmt.Fprintln(outputWriter, "airplane fly")
}

type fighter struct {
}

func (f *fighter) takeOff() {
	fmt.Fprintln(outputWriter, "fighter take off")
}

func (f *fighter) fly() {
	fmt.Fprintln(outputWriter, "fighter fly")
}

type harrier struct {
}

func (h *harrier) takeOff() {
	fmt.Fprintln(outputWriter, "harrier take off")
}

func (h *harrier) fly() {
	fmt.Fprintln(outputWriter, "harrier fly")
}
