package simple_factory

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

// Factory represents product factory
type Factory struct {
}

// CreateProduct creates product
func (f Factory) CreateProduct(arg string) Product {
	if arg == "A" {
		return ConcreteProductA{}
	} else if arg == "B" {
		return ConcreteProductB{}
	}
	return nil
}
