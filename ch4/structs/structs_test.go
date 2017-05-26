package structs

import (
	"strings"
	"testing"
)

func TestStudentInfo(t *testing.T) {
	student := Student{
		gpa:  4.0,
		name: "John Rambo",
		year: 4,
	}
	actual := student.studentInfo()
	if !strings.Contains(actual, "GPA") {
		t.Errorf("%s should have GPA in the string", actual)
	}
}
