package command

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type boardScreen struct {
}

func (bs *boardScreen) open() {
	fmt.Fprintln(outputWriter, "open board screen")
}

func (bs *boardScreen) create() {
	fmt.Fprintln(outputWriter, "create board screen")
}

func (bs *boardScreen) edit() {
	fmt.Fprintln(outputWriter, "edit board screen")
}

type bulletinBoard struct {
	m *menu
}

type menu struct {
	menuItems map[string]*menuItem
}

func newMenu() *menu {
	return &menu{
		menuItems: make(map[string]*menuItem, 0),
	}
}

func (m *menu) addMenuItem(menuName string, aMenuItem *menuItem) {
	m.menuItems[menuName] = aMenuItem
}

func (m *menu) click(menuName string) {
	m.menuItems[menuName].click()
}

type menuItem struct {
	mc menuCommand
}

func newMenuItem(mc menuCommand) *menuItem {
	return &menuItem{
		mc: mc,
	}
}

func (m *menuItem) click() {
	m.mc.execute()
}

type menuCommand interface {
	execute()
}

type defaultMenuCommand struct {
	bs *boardScreen
}

type openMenuCommand struct {
	defaultMenuCommand
}

func (o *openMenuCommand) execute() {
	o.bs.open()
}

type createMenuCommand struct {
	defaultMenuCommand
}

func (c *createMenuCommand) execute() {
	c.bs.create()
}

type editMenuCommand struct {
	defaultMenuCommand
}

func (e *editMenuCommand) execute() {
	e.bs.edit()
}
