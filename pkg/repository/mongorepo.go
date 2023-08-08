package repository

import (
	"context"
	"errors"

	"students/app/models"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBStudentRepository represents a MongoDB implementation of the StudentRepository
type MongoDBStudentRepository struct {
	collection *mongo.Collection
}

// NewMongoDBStudentRepository creates a new instance of MongoDBStudentRepository
func NewMongoDBStudentRepository(client *mongo.Client) *MongoDBStudentRepository {
	return &MongoDBStudentRepository{
		collection: client.Database("students").Collection("collection"),
	}
}

// GetStudents retrieves all students from the MongoDB repository
func (r *MongoDBStudentRepository) GetStudents() ([]models.Student, error) {
	// MongoDB-specific implementation to fetch students
	// ...
	ctx := context.TODO()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []models.Student
	for cursor.Next(ctx) {
		var student models.Student
		err := cursor.Decode(&student)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

// GetStudentByID retrieves a student by ID from the MongoDB repository
func (r *MongoDBStudentRepository) GetStudentByID(id string) (*models.Student, error) {
	// MongoDB-specific implementation to fetch a student by ID
	// ...
	// Parse the given string ID into an ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	filter := bson.M{"_id": objID}

	var student models.Student
	err = r.collection.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &student, nil
}

// CreateStudent adds a new student to the MongoDB repository
func (r *MongoDBStudentRepository) CreateStudent(student *models.Student) error {
	// MongoDB-specific implementation to create a student
	// ...
	ctx := context.TODO()

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Replace the plain-text password with the hashed password
	student.Password = string(hashedPassword)

	// Store the updated student object in the MongoDB collection
	_, err = r.collection.InsertOne(ctx, student)
	if err != nil {
		return err
	}

	return nil
}

// UpdateStudent updates a student in the MongoDB repository
// func (r *MongoDBStudentRepository) UpdateStudent(student *models.Student) error {
// 	// MongoDB-specific implementation to update a student
// 	// ...
// 	ctx := context.TODO()

// 	filter := bson.M{"id": student.ID}

// 	update := bson.M{
// 		"$set": bson.M{
// 			"name":   student.Name,
// 			"id":     student.ID,
// 			"roolno": student.RollNo,
// 			"class":  student.Class,
// 		},
// 	}

// 	_, err := r.collection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// UpdateStudent updates a student in the MongoDB repository
func (r *MongoDBStudentRepository) UpdateStudent(student *models.Student) error {
	filter := bson.M{"_id": student.MongoID}
	update := bson.M{"$set": bson.M{
		"name":     student.Name,
		"rollno":   student.RollNo,
		"class":    student.Class,
	}}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}


// DeleteStudent deletes a student by ID from the MongoDB repository
func (r *MongoDBStudentRepository) DeleteStudent(id string) error {
	// MongoDB-specific implementation to delete a student
	// ...
	ctx := context.TODO()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	_, err = r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}