package mediator

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestWindow(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	w := newWindow()
	tp := &textPane{}
	tp.setPaneMediator(w)
	lp := &listPane{}
	lp.setPaneMediator(w)
	gp := &graphicPane{}
	gp.setPaneMediator(w)

	w.tp = tp
	w.lp = lp
	w.gp = gp

	tp.changed()
	assert.Equal(t, "text pane changed\nlist pane update\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	lp.changed()
	assert.Equal(t, "list pane changed\ntext pane update\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	gp.changed()
	assert.Equal(t, "graphic pane changed\ntext pane update\nlist pane update\n", outputWriter.(*strings.Builder).String())
}
