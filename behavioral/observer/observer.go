package observer

type observer interface {
	update() string
}

type observable interface {
	attach(obs observer)
	detach(obs observer)
	notify() string
}

type subject struct {
	observers []observer
}

func newSubject() *subject {
	return &subject{
		observers: make([]observer, 0),
	}
}

func (s *subject) attach(obs observer) {
	s.observers = append(s.observers, obs)
}

func (s *subject) detach(obs observer) {}

func (s *subject) notify() string {
	result := ""
	for _, obs := range s.observers {
		result += obs.update() + " "
	}
	return result[:len(result)-1]
}

type concreteObserverA struct {
}

func (obs concreteObserverA) update() string {
	return "concreteObserverA"
}

type concreteObserverB struct {
}

func (obs concreteObserverB) update() string {
	return "concreteObserverB"
}

type concreteObserverC struct {
}

func (obs concreteObserverC) update() string {
	return "concreteObserverC"
}
