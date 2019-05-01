package prototype

type Prototype interface {
	Name() string
	Clone() Prototype
}

type ConcretePrototypeA struct {
	name string
}

func (p ConcretePrototypeA) Clone() Prototype {
	n := p
	return n
}

func (p ConcretePrototypeA) Name() string {
	return p.name
}
