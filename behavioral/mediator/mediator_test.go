package mediator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	mediator := newConcreteMediator()
	colleagueA := newConcreteColleagueA()
	colleagueB := newConcreteColleagueB()
	colleagueC := newConcreteColleagueC()

	colleagueA.setMediator(mediator)
	colleagueB.setMediator(mediator)
	colleagueC.setMediator(mediator)

	mediator.colleagueA = colleagueA
	mediator.colleagueB = colleagueB
	mediator.colleagueC = colleagueC

	assert.Equal(t, colleagueA.operation(), "colleagueB_methodB")
	assert.Equal(t, colleagueB.operation(), "colleagueA_methodA colleagueC_methodC")
	assert.Equal(t, colleagueC.operation(), "colleagueA_methodA")
}
