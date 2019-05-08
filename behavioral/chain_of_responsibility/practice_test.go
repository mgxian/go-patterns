package chain_of_responsibility

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeaveRequests(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	aDirector := newDirector("will")
	aManager := newManager("mao")
	aGeneralManager := newGeneralManager("max")

	aDirector.setNextApprover(aManager)
	aManager.setNextApprover(aGeneralManager)

	var lr *leaveRequest
	lr = newLeaveleaveRequest("mgxian", 2)
	aDirector.handleLeaveRequest(lr)
	assert.Equal(t, "director will approved mgxian's 2 days leave\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	lr = newLeaveleaveRequest("mgxian", 5)
	aDirector.handleLeaveRequest(lr)
	assert.Equal(t, "manager mao approved mgxian's 5 days leave\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	lr = newLeaveleaveRequest("mgxian", 20)
	aDirector.handleLeaveRequest(lr)
	assert.Equal(t, "generalManager max approved mgxian's 20 days leave\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	lr = newLeaveleaveRequest("mgxian", 40)
	aDirector.handleLeaveRequest(lr)
	assert.Equal(t, "nobody can approve mgxian's 40 days leave\n", outputWriter.(*strings.Builder).String())
}
