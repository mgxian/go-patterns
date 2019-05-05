package abstract_factory

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type operationController interface {
	show()
}

type symbianOperationController struct {
}

func newSymbianOperationController() *symbianOperationController {
	return &symbianOperationController{}
}

func (s *symbianOperationController) show() {
	fmt.Fprintln(outputWriter, "I'm symbian operation controller")
}

type androidOperationController struct {
}

func newAndroidOperationController() *androidOperationController {
	return &androidOperationController{}
}

func (s *androidOperationController) show() {
	fmt.Fprintln(outputWriter, "I'm android operation controller")
}

type windowsMobileOperationController struct {
}

func newWindowsMobileOperationController() *windowsMobileOperationController {
	return &windowsMobileOperationController{}
}

func (s *windowsMobileOperationController) show() {
	fmt.Fprintln(outputWriter, "I'm windows mobile operation controller")
}

type interfaceController interface {
	show()
}

type symbianInterfaceController struct {
}

func newSymbianInterfaceController() *symbianInterfaceController {
	return &symbianInterfaceController{}
}

func (s *symbianInterfaceController) show() {
	fmt.Fprintln(outputWriter, "I'm symbian interface controller")
}

type androidInterfaceController struct {
}

func newAndroidInterfaceController() *androidInterfaceController {
	return &androidInterfaceController{}
}

func (s *androidInterfaceController) show() {
	fmt.Fprintln(outputWriter, "I'm android interface controller")
}

type windowsMobileInterfaceController struct {
}

func newWindowsMobileInterfaceController() *windowsMobileInterfaceController {
	return &windowsMobileInterfaceController{}
}

func (s *windowsMobileInterfaceController) show() {
	fmt.Fprintln(outputWriter, "I'm windows mobile interface controller")
}

type mobileGameFactory interface {
	createOperationController() operationController
	createInterfaceController() interfaceController
}

type symbianFactory struct {
}

func newSymbianFactory() *symbianFactory {
	return &symbianFactory{}
}

func (s *symbianFactory) createOperationController() operationController {
	return newSymbianOperationController()
}

func (s *symbianFactory) createInterfaceController() interfaceController {
	return newSymbianInterfaceController()
}

type androidFactory struct {
}

func newAndroidFactory() *androidFactory {
	return &androidFactory{}
}

func (s *androidFactory) createOperationController() operationController {
	return newAndroidOperationController()
}

func (s *androidFactory) createInterfaceController() interfaceController {
	return newAndroidInterfaceController()
}

type windowsMobileFactory struct {
}

func newWindowsMobileFactory() *windowsMobileFactory {
	return &windowsMobileFactory{}
}

func (s *windowsMobileFactory) createOperationController() operationController {
	return newWindowsMobileOperationController()
}

func (s *windowsMobileFactory) createInterfaceController() interfaceController {
	return newWindowsMobileInterfaceController()
}
