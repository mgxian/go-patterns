package interpreter

import (
	"container/list"
	"strings"
)

type expression interface {
	interpret() int
}

type number struct {
	integer int
}

func (n *number) interpret() int {
	return n.integer
}

type plus struct {
	left  expression
	right expression
}

func (p *plus) interpret() int {
	return p.left.interpret() + p.right.interpret()
}

type minus struct {
	left  expression
	right expression
}

func (m *minus) interpret() int {
	return m.left.interpret() - m.right.interpret()
}

type variable struct {
	name      string
	variables map[string]expression
}

func (v *variable) interpret() int {
	value, found := v.variables[v.name]
	if found == false {
		return 0
	}
	return value.interpret()
}

type evaluator struct {
	syntaxTree expression
}

func newEvaluator(aExpression string, variables map[string]expression) *evaluator {
	expressionStack := newStack()
	for _, token := range strings.Split(aExpression, " ") {
		if token == "+" {
			right := expressionStack.pop().(expression)
			left := expressionStack.pop().(expression)
			subExpression := &plus{left, right}
			expressionStack.push(subExpression)
		} else if token == "-" {
			right := expressionStack.pop().(expression)
			left := expressionStack.pop().(expression)
			subExpression := &minus{left, right}
			expressionStack.push(subExpression)
		} else {
			expressionStack.push(&variable{token, variables})
		}
	}
	syntaxTree := expressionStack.pop().(expression)
	return &evaluator{syntaxTree}
}

func (e *evaluator) execute() int {
	return e.syntaxTree.interpret()
}

type stack struct {
	l *list.List
}

func newStack() *stack {
	return &stack{
		l: list.New(),
	}
}

func (s *stack) push(v interface{}) {
	s.l.PushBack(v)
}

func (s *stack) pop() interface{} {
	e := s.l.Back()
	return s.l.Remove(e)
}
