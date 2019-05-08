package command

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBulletinBoard(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	m := newMenu()
	bb := bulletinBoard{m}
	bb.m.addMenuItem("open", newMenuItem(new(openMenuCommand)))
	bb.m.addMenuItem("create", newMenuItem(new(createMenuCommand)))
	bb.m.addMenuItem("edit", newMenuItem(new(editMenuCommand)))

	bb.m.click("open")
	assert.Equal(t, "open board screen\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	bb.m.click("create")
	assert.Equal(t, "create board screen\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	bb.m.click("edit")
	assert.Equal(t, "edit board screen\n", outputWriter.(*strings.Builder).String())
}
