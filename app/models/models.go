// models.go

package models

//This is for Mysql
// Student represents the student entity
// type Student struct {
// 	ID               uint     `json:"id" gorm:"primarykey"`
// 	Name             string   `json:"name"`
// 	RollNo           string   `json:"rollno"`
// 	Class            string   `json:"class"`
// 	OptionalSubjects []string `json:"optional_subjects" gorm:"type:json"`
// 	StudentRank      int      `json:"student_rank"` // Renamed from "Rank" to "StudentRank"
// }

// This is working for mongoDB
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

// This is working for postgress sql
// type Student struct {
// 	ID               string   `json:"id" gorm:"primarykey"`
// 	Name             string   `json:"name"`
// 	RollNo           string   `json:"rollno"`
// 	Class            string   `json:"class"`
// 	OptionalSubjects []string `json:"optional_subjects" gorm:"type:text[]"` // I'm getting error here
// 	Rank             int      `json:"student_rank"`
// }
