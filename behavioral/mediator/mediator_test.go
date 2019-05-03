package mediator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMediator(t *testing.T) {
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

	assert.Equal(t, "colleagueB_methodB", colleagueA.operation())
	assert.Equal(t, "colleagueA_methodA colleagueC_methodC", colleagueB.operation())
	assert.Equal(t, "colleagueA_methodA", colleagueC.operation())
}
