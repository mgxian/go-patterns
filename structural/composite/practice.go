package composite

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type uiComponent interface {
	click()
	add(uiComponent)
	remove(uiComponent)
	getChildren() []uiComponent
}

type unitUIComponent struct{}

func (u *unitUIComponent) add(uic uiComponent)    {}
func (u *unitUIComponent) remove(uic uiComponent) {}
func (u *unitUIComponent) getChildren() []uiComponent {
	return []uiComponent{}
}

type buttonUIComponent struct {
	unitUIComponent
	name string
}

func newButtonUIComponent(name string) *buttonUIComponent {
	return &buttonUIComponent{
		name: name,
	}
}

func (b *buttonUIComponent) click() {
	fmt.Fprintln(outputWriter, b.name+" button is clicked")
}

type textboxUIComponent struct {
	unitUIComponent
	name string
}

func newTextboxUIComponent(name string) *textboxUIComponent {
	return &textboxUIComponent{
		name: name,
	}
}

func (t *textboxUIComponent) click() {
	fmt.Fprintln(outputWriter, t.name+" textbox is clicked")
}

type containerUIComponent struct {
	children []uiComponent
}

func (c *containerUIComponent) add(uic uiComponent) {
	c.children = append(c.children, uic)
}
func (c *containerUIComponent) remove(uic uiComponent) {}
func (c *containerUIComponent) getChildren() []uiComponent {
	return c.children
}

type windowUIComponent struct {
	containerUIComponent
	name string
}

func newWindowUIComponent(name string) *windowUIComponent {
	return &windowUIComponent{
		name: name,
	}
}

func (w *windowUIComponent) click() {
	fmt.Fprintln(outputWriter, w.name+" window is clicked and propagate to children")
	for _, child := range w.children {
		child.click()
	}
}

type middlePanelUIComponent struct {
	containerUIComponent
	name string
}

func newMiddlePanelUIComponent(name string) *middlePanelUIComponent {
	return &middlePanelUIComponent{
		name: name,
	}
}

func (m *middlePanelUIComponent) click() {
	fmt.Fprintln(outputWriter, m.name+" middle panel is clicked and propagate to children")
	for _, child := range m.children {
		child.click()
	}
}
