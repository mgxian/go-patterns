package prototype

import (
	"testing"

	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestCustomerClone(t *testing.T) {
	aAdress := newAddress("shangHai", "zhendan")
	aCustomer := newCustomer("mgxian", aAdress)
	bCustomer := aCustomer.clone().(*customer)

	assert.Equal(t, aCustomer.name, bCustomer.name)

	bCustomer.name = "will"
	assert.NotEqual(t, bCustomer.name, aCustomer.name)
	assert.Equal(t, unsafe.Pointer(aCustomer.aAdress), unsafe.Pointer(bCustomer.aAdress))

	bCustomer.aAdress.street = "station"
	assert.Equal(t, aCustomer.aAdress.street, bCustomer.aAdress.street)
}

func TestCustomerDeepClone(t *testing.T) {
	aAdress := newAddress("shangHai", "zhendan")
	aCustomer := newCustomer("mgxian", aAdress)
	bCustomer := aCustomer.deepClone().(*customer)

	assert.Equal(t, aCustomer.name, bCustomer.name)

	bCustomer.name = "will"
	assert.NotEqual(t, bCustomer.name, aCustomer.name)
	assert.NotEqual(t, unsafe.Pointer(aCustomer.aAdress), unsafe.Pointer(bCustomer.aAdress))

	bCustomer.aAdress.street = "station"
	assert.NotEqual(t, aCustomer.aAdress.street, bCustomer.aAdress.street)
}
