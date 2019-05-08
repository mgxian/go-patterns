package mediator

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type paneType int

const (
	text paneType = iota
	list
	graphic
)

type window struct {
	tp *textPane
	lp *listPane
	gp *graphicPane
}

func newWindow() *window {
	return &window{}
}

func (w *window) changed(pt paneType) {
	switch pt {
	case text:
		w.lp.update()
	case list:
		w.tp.update()
	case graphic:
		w.tp.update()
		w.lp.update()
	}
}

type pane struct {
	w *window
}

func (p *pane) setPaneMediator(w *window) {
	p.w = w
}

type textPane struct {
	pane
}

func (tp *textPane) update() {
	fmt.Fprintln(outputWriter, "text pane update")
}

func (tp *textPane) changed() {
	fmt.Fprintln(outputWriter, "text pane changed")
	tp.w.changed(text)
}

type listPane struct {
	pane
}

func (lp *listPane) update() {
	fmt.Fprintln(outputWriter, "list pane update")
}

func (lp *listPane) changed() {
	fmt.Fprintln(outputWriter, "list pane changed")
	lp.w.changed(list)
}

type graphicPane struct {
	pane
}

func (gp *graphicPane) update() {
	fmt.Fprintln(outputWriter, "graphic pane update")
}

func (gp *graphicPane) changed() {
	fmt.Fprintln(outputWriter, "graphic pane changed")
	gp.w.changed(graphic)
}
