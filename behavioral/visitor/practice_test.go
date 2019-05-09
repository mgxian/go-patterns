package visitor

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAwardCheck(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	t1 := newTeacher("t1", 5, 70)
	t2 := newTeacher("t2", 11, 70)
	t3 := newTeacher("t3", 8, 85)
	t4 := newTeacher("t4", 12, 88)

	s1 := newStudent("s1", 1, 78)
	s2 := newStudent("s2", 4, 68)
	s3 := newStudent("s3", 1, 90)
	s4 := newStudent("s4", 3, 86)

	var acceptors []acceptor
	acceptors = append(acceptors, t1, s1, s2, t2, t3, s3, s4, t4)

	acc := newAwardCheckController(10, 2, 80, 80)

	for _, a := range acceptors {
		a.accept(acc)
	}

	expected := "student s2 get scientific research award\n" +
		"teacher t2 get scientific research award\n" +
		"teacher t3 get outstanding academic award\n" +
		"student s3 get outstanding academic award\n" +
		"student s4 get scientific research award\n" +
		"student s4 get outstanding academic award\n" +
		"teacher t4 get scientific research award\n" +
		"teacher t4 get outstanding academic award\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
}
