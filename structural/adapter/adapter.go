package adapter

type target interface {
	request() string
}

type adaptee struct {
}

func newAdaptee() adaptee {
	return adaptee{}
}

func (a adaptee) specificRequest() string {
	return "raw response"
}

type adapter struct {
	aAdaptee adaptee
}

func newAdapter(aAdaptee adaptee) *adapter {
	return &adapter{
		aAdaptee: aAdaptee,
	}
}

func (a adapter) request() string {
	rawResponse := a.aAdaptee.specificRequest()
	return "adapter response + " + rawResponse
}
