package composite

type component interface {
	add(c component)
	remove(c component)
	getChildren() []component
	operation() string
}

type leaf struct {
	name string
}

func newLeaf(name string) *leaf {
	return &leaf{
		name: name,
	}
}

func (l leaf) add(c component)          {}
func (l leaf) remove(c component)       {}
func (l leaf) getChildren() []component { return nil }

func (l leaf) operation() string {
	return l.name
}

type composite struct {
	children []component
}

func newComposite() *composite {
	return &composite{
		children: make([]component, 0),
	}
}

func (c *composite) add(aComponent component) {
	c.children = append(c.children, aComponent)
}

func (c *composite) remove(aComponent component) {}

func (c *composite) getChildren() []component {
	return c.children
}

func (c *composite) operation() string {
	result := ""
	for _, leaf := range c.children {
		result += leaf.operation()
	}
	return result
}
