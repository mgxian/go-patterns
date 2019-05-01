package abstract_factory

// ProductA describes abstract product
type ProductA interface {
	show() string
}

// ConcreteProductA1 represents concrete product of A
type ConcreteProductA1 struct {
}

func (p ConcreteProductA1) show() string {
	return "ProductA1"
}

// ConcreteProductA2 represents concrete product of A
type ConcreteProductA2 struct {
}

func (p ConcreteProductA2) show() string {
	return "ProductA2"
}

// ProductB describes abstract product
type ProductB interface {
	show() string
}

// ConcreteProductB1 represents concrete product of B
type ConcreteProductB1 struct {
}

func (p ConcreteProductB1) show() string {
	return "ProductB1"
}

// ConcreteProductB2 represents concrete product of B
type ConcreteProductB2 struct {
}

func (p ConcreteProductB2) show() string {
	return "ProductB2"
}

// AbstractFactory describe abstract factory.
type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

// ConcreteFactory1 represents concrete product factory of A
type ConcreteFactory1 struct {
}

// CreateProductA creates product A
func (f ConcreteFactory1) CreateProductA() ProductA {
	return ConcreteProductA1{}
}

// CreateProductB creates product B
func (f ConcreteFactory1) CreateProductB() ProductB {
	return ConcreteProductB1{}
}

// ConcreteFactory2 represents concrete product factory of A
type ConcreteFactory2 struct {
}

// CreateProductA creates product A
func (f ConcreteFactory2) CreateProductA() ProductA {
	return ConcreteProductA2{}
}

// CreateProductB creates product B
func (f ConcreteFactory2) CreateProductB() ProductB {
	return ConcreteProductB2{}
}
