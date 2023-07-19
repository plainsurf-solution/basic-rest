// models.go

package models

// Student represents the student entity
type Student struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	RollNo         string   `json:"rollno"`
	Class          string   `json:"class"`
	OptionalSubjec []string `json:"optional_subjects"`
	Rank           int      `json:"rank"`
}
