package facade

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type mobilePhoneManager struct {
	c *contacts
	m *messages
	p *pictures
	s *songs
}

func newMobilePhoneManager() *mobilePhoneManager {
	return &mobilePhoneManager{
		c: &contacts{},
		m: &messages{},
		p: &pictures{},
		s: &songs{},
	}
}

func (m *mobilePhoneManager) oneKeyBackupToSDCard() {
	fmt.Fprintln(outputWriter, "one key backup to sd card")
	m.c.backupToSDCard()
	m.m.backupToSDCard()
	m.p.backupToSDCard()
	m.s.backupToSDCard()
}

type contacts struct{}

func (c *contacts) backupToSDCard() {
	fmt.Fprintln(outputWriter, "contacts backup to sd card")
}

type messages struct{}

func (c *messages) backupToSDCard() {
	fmt.Fprintln(outputWriter, "messages backup to sd card")
}

type pictures struct{}

func (c *pictures) backupToSDCard() {
	fmt.Fprintln(outputWriter, "pictures backup to sd card")
}

type songs struct{}

func (c *songs) backupToSDCard() {
	fmt.Fprintln(outputWriter, "songs backup to sd card")
}
