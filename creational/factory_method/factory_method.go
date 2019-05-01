package factory_method

// Product describes abstract product
type Product interface {
	show() string
}

// ConcreteProductA represents concrete product of A
type ConcreteProductA struct {
}

func (p ConcreteProductA) show() string {
	return "ProductA"
}

// ConcreteProductB represents concrete product of B
type ConcreteProductB struct {
}

func (p ConcreteProductB) show() string {
	return "ProductB"
}

// Factory describe abstract factory.
type Factory interface {
	CreateProduct() Product
}

// ConcreteAFactory represents concrete product factory of A
type ConcreteAFactory struct {
}

// CreateProduct creates product
func (f ConcreteAFactory) CreateProduct() Product {
	return ConcreteProductA{}
}

// ConcreteBFactory represents concrete product factory of B
type ConcreteBFactory struct {
}

// CreateProduct creates product
func (f ConcreteBFactory) CreateProduct() Product {
	return ConcreteProductB{}
}
