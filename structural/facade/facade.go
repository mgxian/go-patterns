package facade

type facade struct {
	aSubSystemA subSystemA
	aSubSystemB subSystemB
	aSubSystemC subSystemC
}

func newFacade(aSubSystemA subSystemA, aSubSystemB subSystemB, aSubSystemC subSystemC) *facade {
	return &facade{
		aSubSystemA: aSubSystemA,
		aSubSystemB: aSubSystemB,
		aSubSystemC: aSubSystemC,
	}
}

func (f facade) operation() string {
	return f.aSubSystemA.operation() + f.aSubSystemB.operation() + f.aSubSystemC.operation()
}

type subSystemA struct {
}

func (s subSystemA) operation() string {
	return "A"
}

type subSystemB struct {
}

func (s subSystemB) operation() string {
	return "B"
}

type subSystemC struct {
}

func (s subSystemC) operation() string {
	return "C"
}
