package state

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type pointType int

// level points
const (
	PrimaryPoints      pointType = 0
	SecondaryPoints    pointType = 100
	ProfessionalPoints pointType = 500
	FinalPoints        pointType = 1000
)

type level interface {
	play()
	doubleScore()
	changeCards()
	peekCards()
}

type primary struct {
}

func (p *primary) play() {
	fmt.Fprintln(outputWriter, "I can play")
}

func (p *primary) doubleScore() {
	fmt.Fprintln(outputWriter, "I can not doubleScore")
}

func (p *primary) changeCards() {
	fmt.Fprintln(outputWriter, "I can not changeCards")
}

func (p *primary) peekCards() {
	fmt.Fprintln(outputWriter, "I can not peekCards")
}

type secondary struct {
	primary
}

func (s *secondary) doubleScore() {
	fmt.Fprintln(outputWriter, "I can doubleScore")
}

type professional struct {
	secondary
}

func (p *professional) changeCards() {
	fmt.Fprintln(outputWriter, "I can changeCards")
}

type final struct {
	professional
}

func (f *final) peekCards() {
	fmt.Fprintln(outputWriter, "I can peekCards")
}

type cardGamePlayer struct {
	level
	name   string
	points pointType
}

func newCardGamePlayer(name string) *cardGamePlayer {
	return &cardGamePlayer{
		name:   name,
		points: 0,
		level:  &primary{},
	}
}

func (c *cardGamePlayer) changeLevel() {
	if c.points >= FinalPoints {
		c.level = &final{}
		return
	}

	if c.points >= ProfessionalPoints {
		c.level = &professional{}
		return
	}

	if c.points >= SecondaryPoints {
		c.level = &secondary{}
		return
	}

	c.level = &primary{}
}

func (c *cardGamePlayer) addPoints(p pointType) {
	c.points += p
	c.changeLevel()
}

func (c *cardGamePlayer) deductPoints(p pointType) {
	c.points -= p
	c.changeLevel()
}
