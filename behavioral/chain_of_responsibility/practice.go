package chain_of_responsibility

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type leaveRequest struct {
	name string
	days int
}

func newLeaveleaveRequest(name string, days int) *leaveRequest {
	return &leaveRequest{
		name: name,
		days: days,
	}
}

type leaveApprover interface {
	handleLeaveRequest(lr *leaveRequest)
	setNextApprover(la leaveApprover)
}

type defaultApprover struct {
	nextApprover leaveApprover
}

func (da *defaultApprover) setNextApprover(la leaveApprover) {
	da.nextApprover = la
}

type director struct {
	defaultApprover
	name string
}

func newDirector(name string) *director {
	return &director{
		name: name,
	}
}

func (d *director) handleLeaveRequest(lr *leaveRequest) {
	if lr.days < 3 {
		fmt.Fprintf(outputWriter, "director %s approved %s's %d days leave\n", d.name, lr.name, lr.days)
		return
	}
	d.nextApprover.handleLeaveRequest(lr)
}

type manager struct {
	defaultApprover
	name string
}

func newManager(name string) *manager {
	return &manager{
		name: name,
	}
}

func (m *manager) handleLeaveRequest(lr *leaveRequest) {
	if lr.days >= 3 && lr.days < 10 {
		fmt.Fprintf(outputWriter, "manager %s approved %s's %d days leave\n", m.name, lr.name, lr.days)
		return
	}
	m.nextApprover.handleLeaveRequest(lr)
}

type generalManager struct {
	defaultApprover
	name string
}

func newGeneralManager(name string) *generalManager {
	return &generalManager{
		name: name,
	}
}

func (g *generalManager) handleLeaveRequest(lr *leaveRequest) {
	if lr.days >= 10 && lr.days < 30 {
		fmt.Fprintf(outputWriter, "generalManager %s approved %s's %d days leave\n", g.name, lr.name, lr.days)
		return
	}
	fmt.Fprintf(outputWriter, "nobody can approve %s's %d days leave\n", lr.name, lr.days)
}
