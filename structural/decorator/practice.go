package decorator

import (
	"strings"
)

type crypto interface {
	encrypt(string) string
	decrypt(string) string
}

type simpleShiftCrypto struct {
	shiftNumber int
}

func newSimpleShiftCrypto(n int) *simpleShiftCrypto {
	return &simpleShiftCrypto{
		shiftNumber: n,
	}
}

func (s *simpleShiftCrypto) encrypt(clearData string) string {
	result := new(strings.Builder)
	for _, c := range clearData {
		result.WriteByte(byte((int(c) + s.shiftNumber) % 128))
	}
	return result.String()
}

func (s *simpleShiftCrypto) decrypt(cipherData string) string {
	result := new(strings.Builder)
	for _, c := range cipherData {
		n := int(c) - s.shiftNumber
		if n < 0 {
			n += 128
		}
		result.WriteByte(byte(n))
	}
	return result.String()
}

type strongCrypto struct {
	c crypto
}

func (s *strongCrypto) encrypt(clearData string) string  { return "" }
func (s *strongCrypto) decrypt(cipherData string) string { return "" }

type reverseCrypto struct {
	strongCrypto
}

func newReverseCrypto(c crypto) *reverseCrypto {
	return &reverseCrypto{
		strongCrypto{c: c},
	}
}

func (r *reverseCrypto) reverse(aString string) string {
	result := new(strings.Builder)
	for i := len(aString) - 1; i >= 0; i-- {
		result.WriteByte(aString[i])
	}
	return result.String()
}

func (r *reverseCrypto) encrypt(clearData string) string {
	return r.reverse(r.c.encrypt(clearData))
}

func (r *reverseCrypto) decrypt(cipherData string) string {
	return r.c.decrypt(r.reverse((cipherData)))
}

type aesFakeCrypto struct {
	strongCrypto
	key []byte
}

func newAESCrypto(key string, c crypto) *aesFakeCrypto {
	return &aesFakeCrypto{
		key:          []byte(key),
		strongCrypto: strongCrypto{c: c},
	}
}

func (m *aesFakeCrypto) encrypt(clearData string) string {

	return "#*#" + m.c.encrypt(clearData) + "#*#"
}

func (m *aesFakeCrypto) decrypt(cipherData string) string {
	return m.c.decrypt(cipherData[3 : len(cipherData)-3])
}
