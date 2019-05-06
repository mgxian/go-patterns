package adapter

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestMySQLUserRepository(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var ur userRepository = new(mysqlUserRepository)
	u := ur.getUserByName("mgxian")
	assert.Equal(t, newUser("mgxian", "123456", "mgxian@will.com"), u)

	u.name = "will"
	ur.saveUser(u)
	assert.Equal(t, "will\n123456\nmgxian@will.com\n", outputWriter.(*strings.Builder).String())
}

func TestSecureMySQLUserRepository(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var ur userRepository = new(mysqlSecureUserRepository)
	u := ur.getUserByName("mgxian")
	assert.Equal(t, newUser("mgxian", "123456", "mgxian@will.com"), u)

	u.name = "will"
	ur.saveUser(u)
	assert.Equal(t, "will\nencrypt[123456]encrypt\nencrypt[mgxian@will.com]encrypt\n", outputWriter.(*strings.Builder).String())
}

func TestSecureMySQLUserRepositoryUseComposite(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var ur userRepository = new(mysqlSecureUserRepositoryUseComposite)
	u := ur.getUserByName("mgxian")
	assert.Equal(t, newUser("mgxian", "123456", "mgxian@will.com"), u)

	u.name = "will"
	ur.saveUser(u)
	assert.Equal(t, "will\nencrypt[123456]encrypt\nencrypt[mgxian@will.com]encrypt\n", outputWriter.(*strings.Builder).String())
}
