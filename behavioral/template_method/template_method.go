package template_method

type template interface {
	operationA() string
	operationB() string
	operationC() string
}

type abstractClass struct {
	template
}

func newAbstractClass(t template) *abstractClass {
	return &abstractClass{
		template: t,
	}
}

func (ac abstractClass) templateMethod() string {
	result := ac.template.operationA()
	result += " " + ac.template.operationB()
	result += " " + ac.template.operationC()
	return result
}

type concreteClass struct {
}

func newConcreteClass() *concreteClass {
	return &concreteClass{}
}

func (c concreteClass) operationA() string {
	return "operationA"
}

func (c concreteClass) operationB() string {
	return "operationB"
}

func (c concreteClass) operationC() string {
	return "operationC"
}
