package visitor

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type awardChecker interface {
	teacherScientificResearchCheck(t *teacher)
	teacheroutstandingAcademicCheck(t *teacher)

	studentScientificResearchCheck(s *student)
	studentoutstandingAcademicCheck(s *student)
}

type acceptor interface {
	accept(ac awardChecker)
}

type awardCheckController struct {
	teacherAwardPaperMin int
	studentAwardPaperMin int
	teacherAwardScoreMin int
	studentAwardScoreMin int
}

func newAwardCheckController(tpm, spm, tsm, ssm int) *awardCheckController {
	return &awardCheckController{
		teacherAwardPaperMin: tpm,
		teacherAwardScoreMin: tsm,
		studentAwardPaperMin: spm,
		studentAwardScoreMin: ssm,
	}
}

func (acc *awardCheckController) teacherScientificResearchCheck(t *teacher) {
	if t.paperCount >= acc.teacherAwardPaperMin {
		fmt.Fprintln(outputWriter, "teacher "+t.name+" get scientific research award")
	}
}

func (acc *awardCheckController) teacheroutstandingAcademicCheck(t *teacher) {
	if t.teachScore >= acc.teacherAwardScoreMin {
		fmt.Fprintln(outputWriter, "teacher "+t.name+" get outstanding academic award")
	}
}

func (acc *awardCheckController) studentScientificResearchCheck(s *student) {
	if s.paperCount >= acc.studentAwardPaperMin {
		fmt.Fprintln(outputWriter, "student "+s.name+" get scientific research award")
	}
}

func (acc *awardCheckController) studentoutstandingAcademicCheck(s *student) {
	if s.studyScore >= acc.studentAwardScoreMin {
		fmt.Fprintln(outputWriter, "student "+s.name+" get outstanding academic award")
	}
}

type teacher struct {
	name       string
	paperCount int
	teachScore int
}

func newTeacher(name string, pc, ts int) *teacher {
	return &teacher{
		name:       name,
		paperCount: pc,
		teachScore: ts,
	}
}

func (t *teacher) accept(ac awardChecker) {
	ac.teacherScientificResearchCheck(t)
	ac.teacheroutstandingAcademicCheck(t)
}

type student struct {
	name       string
	paperCount int
	studyScore int
}

func newStudent(name string, pc, ss int) *student {
	return &student{
		name:       name,
		paperCount: pc,
		studyScore: ss,
	}
}

func (s *student) accept(ac awardChecker) {
	ac.studentScientificResearchCheck(s)
	ac.studentoutstandingAcademicCheck(s)
}
