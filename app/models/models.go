// models.go

package models

type Student struct {
	ID               uint     `json:"id" gorm:"primarykey"`
	Email            string   `json:"email"`
	Password         string   `json:"password"`
	Name             string   `json:"name"`
	RollNo           string   `json:"rollno"`
	Class            string   `json:"class"`
	OptionalSubjects []string `json:"optional_subjects"`
	StudentRank      int      `json:"student_rank"`
}


