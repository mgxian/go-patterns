package prototype

// Prototype describe prototype interface
type Prototype interface {
	Clone() Prototype
}

// ConcretePrototype implements prototype
type ConcretePrototype struct {
}

// Clone return the of copy of this struct
func (p ConcretePrototype) Clone() Prototype {
	return p
}
