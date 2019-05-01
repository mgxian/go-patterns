package flyweight

type flyweight struct {
	intrinsicState string
}

func newFlyweight(state string) *flyweight {
	return &flyweight{
		intrinsicState: state,
	}
}

type flyweightFactory struct {
	flyweights map[string]*flyweight
}

func newFlyweightFactory() *flyweightFactory {
	return &flyweightFactory{
		flyweights: make(map[string]*flyweight, 0),
	}
}

func (f *flyweightFactory) GetFlyweight(state string) *flyweight {
	flyweight, ok := f.flyweights[state]
	if !ok {
		flyweight = newFlyweight(state)
		f.flyweights[state] = flyweight
	}
	return flyweight
}
