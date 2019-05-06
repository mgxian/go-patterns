package composite

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUI(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var mainWindow uiComponent = newWindowUIComponent("hello mgxian")
	var mp uiComponent = newMiddlePanelUIComponent("CPU monitor")
	mainWindow.add(mp)
	mainWindow.add(newButtonUIComponent("confirm"))
	mainWindow.add(newButtonUIComponent("cancle"))

	mp.add(newTextboxUIComponent("hostname"))
	var searchButton uiComponent = newButtonUIComponent("search")
	mp.add(searchButton)

	searchButton.click()
	assert.Equal(t, "search button is clicked\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	mp.click()
	expected := "CPU monitor middle panel is clicked and propagate to children\n" +
		"hostname textbox is clicked\n" +
		"search button is clicked\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
}
