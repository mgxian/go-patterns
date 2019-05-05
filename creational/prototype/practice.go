package prototype

type cloneable interface {
	clone() cloneable
	deepClone() cloneable
}

type address struct {
	province string
	street   string
}

func newAddress(province, street string) *address {
	return &address{
		province: province,
		street:   street,
	}
}

func (a *address) clone() cloneable {
	return newAddress(a.province, a.street)
}

func (a *address) deepClone() cloneable {
	return newAddress(a.province, a.street)
}

type customer struct {
	name    string
	aAdress *address
}

func newCustomer(name string, aAdress *address) *customer {
	return &customer{
		name:    name,
		aAdress: aAdress,
	}
}

func (c *customer) clone() cloneable {
	return newCustomer(c.name, c.aAdress)
}

func (c *customer) deepClone() cloneable {
	aAdress := c.aAdress.deepClone().(*address)
	return newCustomer(c.name, aAdress)
}
