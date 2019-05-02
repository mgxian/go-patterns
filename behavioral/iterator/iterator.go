package iterator

type iterator interface {
	first()
	next()
	hasNext() bool
	currentItem() interface{}
}

type concreteIterator struct {
	aConcreteAggregate *concreteAggregate
	cursor             int
}

func newConcreteIterator(aConcreteAggregate *concreteAggregate) *concreteIterator {
	return &concreteIterator{
		aConcreteAggregate: aConcreteAggregate,
	}
}

func (ci *concreteIterator) first() {
	ci.cursor = 0
}

func (ci *concreteIterator) next() {
	if ci.cursor < ci.aConcreteAggregate.size()-1 {
		ci.cursor++
	}
}

func (ci *concreteIterator) hasNext() bool {
	if ci.cursor < ci.aConcreteAggregate.size()-1 {
		return true
	}
	return false
}

func (ci concreteIterator) currentItem() interface{} {
	return ci.aConcreteAggregate.get(ci.cursor)
}

type aggregate interface {
	createIterator() iterator
}

type concreteAggregate struct {
	data []interface{}
}

func newConcreteAggregate() *concreteAggregate {
	return &concreteAggregate{
		data: make([]interface{}, 0),
	}
}

func (ca *concreteAggregate) createIterator() iterator {
	return newConcreteIterator(ca)
}

func (ca *concreteAggregate) add(v interface{}) {
	ca.data = append(ca.data, v)
}

func (ca *concreteAggregate) get(i int) interface{} {
	if i < ca.size() {
		return ca.data[i]
	}
	return nil
}

func (ca *concreteAggregate) size() int {
	return len(ca.data)
}
