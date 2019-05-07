package facade

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneKeyBackup(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	mbm := newMobilePhoneManager()
	mbm.oneKeyBackupToSDCard()

	expected := "one key backup to sd card\n" +
		"contacts backup to sd card\n" +
		"messages backup to sd card\n" +
		"pictures backup to sd card\n" +
		"songs backup to sd card\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
}
