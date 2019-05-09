package state

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardGame(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	p := newCardGamePlayer("mgxian")
	p.play()
	p.doubleScore()
	p.changeCards()
	p.peekCards()

	expected := "I can play\n" +
		"I can not doubleScore\n" +
		"I can not changeCards\n" +
		"I can not peekCards\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	p.addPoints(SecondaryPoints)
	p.play()
	p.doubleScore()
	p.changeCards()
	p.peekCards()

	expected = "I can play\n" +
		"I can doubleScore\n" +
		"I can not changeCards\n" +
		"I can not peekCards\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	p.addPoints(ProfessionalPoints - SecondaryPoints)
	p.play()
	p.doubleScore()
	p.changeCards()
	p.peekCards()

	expected = "I can play\n" +
		"I can doubleScore\n" +
		"I can changeCards\n" +
		"I can not peekCards\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	p.addPoints(FinalPoints - ProfessionalPoints)
	p.play()
	p.doubleScore()
	p.changeCards()
	p.peekCards()

	expected = "I can play\n" +
		"I can doubleScore\n" +
		"I can changeCards\n" +
		"I can peekCards\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
}
