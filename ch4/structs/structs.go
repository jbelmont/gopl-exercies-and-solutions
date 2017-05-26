package structs

import (
	"strconv"
)

// Student struct has fields gpa name and year
type Student struct {
	gpa  float64
	name string
	year int
}

func (student *Student) studentInfo() string {
	stringGPA := strconv.FormatFloat(student.gpa, 'f', 2, 64)
	stringYear := strconv.Itoa(student.year)
	return "GPA: " + stringGPA + "\nName: " + student.name + "\nYear: " + stringYear
}
