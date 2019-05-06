package adapter

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type user struct {
	name     string
	password string
	email    string
}

func newUser(name, password, email string) user {
	return user{
		name:     name,
		password: password,
		email:    email,
	}
}

type userRepository interface {
	getUserByName(name string) user
	saveUser(u user)
}

type mysqlUserRepository struct {
}

func (m *mysqlUserRepository) getUserByName(name string) user {
	u := user{
		name:     name,
		password: "123456",
		email:    name + "@will.com",
	}
	return u
}

func (m *mysqlUserRepository) saveUser(u user) {
	fmt.Fprintln(outputWriter, u.name)
	fmt.Fprintln(outputWriter, u.password)
	fmt.Fprintln(outputWriter, u.email)
}

type crypto struct {
}

func (c crypto) encrypt(clearData string) string {
	return "encrypt[" + clearData + "]encrypt"
}

func (c crypto) decrypt(cipherData string) string {
	n := len(cipherData)
	return cipherData[8 : n-8]
}

type mysqlSecureUserRepository struct {
	c crypto
}

func (m *mysqlSecureUserRepository) getUserByName(name string) user {
	u := user{
		name:     name,
		password: m.c.decrypt("encrypt[123456]encrypt"),
		email:    m.c.decrypt("encrypt[" + name + "@will.com" + "]encrypt"),
	}
	return u
}

func (m *mysqlSecureUserRepository) saveUser(u user) {
	fmt.Fprintln(outputWriter, u.name)
	fmt.Fprintln(outputWriter, m.c.encrypt(u.password))
	fmt.Fprintln(outputWriter, m.c.encrypt(u.email))
}

type mysqlSecureUserRepositoryUseComposite struct {
	crypto
}

func (m *mysqlSecureUserRepositoryUseComposite) getUserByName(name string) user {
	u := user{
		name:     name,
		password: m.decrypt("encrypt[123456]encrypt"),
		email:    m.decrypt("encrypt[" + name + "@will.com" + "]encrypt"),
	}
	return u
}

func (m *mysqlSecureUserRepositoryUseComposite) saveUser(u user) {
	fmt.Fprintln(outputWriter, u.name)
	fmt.Fprintln(outputWriter, m.encrypt(u.password))
	fmt.Fprintln(outputWriter, m.encrypt(u.email))
}
