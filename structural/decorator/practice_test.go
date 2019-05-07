package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleShiftCrypto(t *testing.T) {
	var c crypto = newSimpleShiftCrypto(10)
	clearText := "mgxian"
	cipherText := c.encrypt(clearText)
	assert.Equal(t, "wq\x02skx", cipherText)
	assert.Equal(t, clearText, c.decrypt(cipherText))
}

func TestSimpleShiftAndReverseCrypto(t *testing.T) {
	var c crypto = newSimpleShiftCrypto(10)
	c = newReverseCrypto(c)
	clearText := "mgxian"
	cipherText := c.encrypt(clearText)
	assert.Equal(t, "xks\x02qw", cipherText)
	assert.Equal(t, clearText, c.decrypt(cipherText))
}

func TestSimpleShiftAndReverseAndAESCrypto(t *testing.T) {
	var c crypto = newSimpleShiftCrypto(10)
	c = newReverseCrypto(c)
	c = newAESCrypto("will", c)
	clearText := "mgxian"
	cipherText := c.encrypt(clearText)
	assert.Equal(t, "#*#xks\x02qw#*#", cipherText)
	assert.Equal(t, clearText, c.decrypt(cipherText))
}
