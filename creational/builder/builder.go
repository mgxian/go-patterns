package builder

// Product represents complex product
type Product struct {
	result string
}

func (p *Product) makePartA() {
	p.result += "A"
}

func (p *Product) makePartB() {
	p.result += "B"
}

func (p *Product) makePartC() {
	p.result += "C"
}

func (p *Product) getResult() string {
	return p.result
}

type builder interface {
	buildPartA()
	buildPartB()
	buildPartC()
	getResult() Product
}

type concreteBuilder struct {
	product Product
}

func newConcreteBuilder() *concreteBuilder {
	return &concreteBuilder{
		product: Product{},
	}
}

func (b *concreteBuilder) buildPartA() {
	b.product.makePartA()
}

func (b *concreteBuilder) buildPartB() {
	b.product.makePartB()
}

func (b *concreteBuilder) buildPartC() {
	b.product.makePartC()
}

func (b *concreteBuilder) getResult() Product {
	return b.product
}

type director struct {
	aBuilder builder
}

func newDirector(aBuilder builder) *director {
	return &director{
		aBuilder: aBuilder,
	}
}

func (d *director) construct() Product {
	d.aBuilder.buildPartA()
	d.aBuilder.buildPartB()
	d.aBuilder.buildPartC()
	return d.aBuilder.getResult()
}
