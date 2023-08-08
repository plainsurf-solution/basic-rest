// models.go

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// type Student struct {
// 	ID               uint     `json:"id" gorm:"primarykey"`
// 	Email            string   `json:"email"`
// 	Password         string   `json:"password"`
// 	Name             string   `json:"name"`
// 	RollNo           string   `json:"rollno"`
// 	Class            string   `json:"class"`
// 	OptionalSubjects []string `json:"optional_subjects"`
// 	StudentRank      int      `json:"student_rank"`
// }
type optional_subjects []string

type Student struct {
	ID               uint     `json:"id" gorm:"primaryKey"`
	MongoID          primitive.ObjectID `json:"mongo_id" bson:"_id,omitempty" gorm:"-"`
	Email            string   `json:"email"`
	Password         string   `json:"password"`
	Name             string   `json:"name"`
	RollNo           string   `json:"rollno"`
	Class            string   `json:"class"`
	OptionalSubjects optional_subjects `json:"optional_subjects" gorm:"serializer:json"`
	StudentRank      int      `json:"student_rank"`
}


