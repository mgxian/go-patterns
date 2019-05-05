package abstract_factory

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSymbianMobileGame(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var mgf mobileGameFactory
	var oc operationController
	var ic interfaceController

	mgf = newSymbianFactory()
	oc = mgf.createOperationController()
	ic = mgf.createInterfaceController()
	oc.show()
	ic.show()

	assert.Equal(t, "I'm symbian operation controller\nI'm symbian interface controller\n",
		outputWriter.(*strings.Builder).String())
}

func TestCreateAndroidMobileGame(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var mgf mobileGameFactory
	var oc operationController
	var ic interfaceController

	mgf = newAndroidFactory()
	oc = mgf.createOperationController()
	ic = mgf.createInterfaceController()
	oc.show()
	ic.show()

	assert.Equal(t, "I'm android operation controller\nI'm android interface controller\n",
		outputWriter.(*strings.Builder).String())
}

func TestCreateWindowsMobileGame(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var mgf mobileGameFactory
	var oc operationController
	var ic interfaceController

	mgf = newWindowsMobileFactory()
	oc = mgf.createOperationController()
	ic = mgf.createInterfaceController()
	oc.show()
	ic.show()

	assert.Equal(t, "I'm windows mobile operation controller\nI'm windows mobile interface controller\n",
		outputWriter.(*strings.Builder).String())
}
