package memento

type originator struct {
	state string
}

func newOriginator(state string) *originator {
	return &originator{
		state: state,
	}
}

func (o originator) getState() string {
	return o.state
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o originator) createMemento() *memento {
	return newMemento(o.state)
}

func (o *originator) restoreMemento(m *memento) {
	o.state = m.getState()
}

type memento struct {
	state string
}

func newMemento(state string) *memento {
	return &memento{
		state: state,
	}
}

func (m memento) getState() string {
	return m.state
}

func (m *memento) setState(state string) {
	m.state = state
}

type caretaker struct {
	m *memento
}

func newCaretaker(m *memento) *caretaker {
	return &caretaker{
		m: m,
	}
}

func (c caretaker) getMemento() *memento {
	return c.m
}

func (c *caretaker) setMemento(m *memento) {
	c.m = m
}
