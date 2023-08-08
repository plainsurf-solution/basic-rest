package repository_test

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"students/app/models"
	"students/pkg/repository"
)

func TestMongoDBStudentRepository(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("students").Collection("collection")
	repo := repository.NewMongoDBStudentRepository(client)

	// Clean up the collection before running tests
	_, err = collection.DeleteMany(context.Background(), nil)
	if err != nil {
		// Check if the error is "document is nil", if so, skip the failure
		if err.Error() != "document is nil" {
			t.Fatalf("Failed to clean up collection: %v", err)
		}
	}

	// Test CreateStudent and GetStudentByID
	student := &models.Student{
		ID:       5,
		MongoID:  primitive.NewObjectID(), // Generate a new ObjectID
		Name:     "John Doe",
		RollNo:   "12345",
		Class:    "10",
		Password: "testpassword",
	}

	err = repo.CreateStudent(student)
	if err != nil {
		t.Fatalf("Failed to create student: %v", err)
	}

	storedStudent, err := repo.GetStudentByID(student.MongoID.Hex()) // Fetch by ObjectID in hexadecimal format
	if err != nil {
		t.Fatalf("Failed to get student by ID: %v", err)
	}
	if storedStudent == nil {
		t.Fatalf("Student with ID %s not found", student.MongoID.Hex())
	}

	if student.Name != storedStudent.Name || student.RollNo != storedStudent.RollNo || student.Class != storedStudent.Class {
		t.Fatalf("Stored student does not match the expected values")
	}

	fmt.Println(storedStudent, "stored student before update")

	// Test UpdateStudent
	// Test UpdateStudent
	originalMongoID := student.MongoID // Store the original MongoID

	student.Name = "updated name"
	student.RollNo = "84"
	student.Class = "10B"
	err = repo.UpdateStudent(student)
	if err != nil {
		t.Fatalf("Failed to update student: %v", err)
	}

	// Fetch the updatedStudent again using the updated MongoID
	updatedStudent, err := repo.GetStudentByID(student.MongoID.Hex()) // Fetch by ObjectID in hexadecimal format
	if err != nil {
		t.Fatalf("Failed to get updated student by ID: %v", err)
	}
	if updatedStudent == nil {
		t.Fatalf("Updated student with ID %s not found", student.MongoID.Hex())
	}

	// Compare with the originalMongoID
	if originalMongoID.Hex() != updatedStudent.MongoID.Hex() {
		t.Fatalf("Updated student's MongoID does not match the expected value")
	}

	// Compare other fields as before
	if student.Name != updatedStudent.Name || student.RollNo != updatedStudent.RollNo || student.Class != updatedStudent.Class {
		fmt.Println(student.Name, updatedStudent.Name, student.RollNo, updatedStudent.RollNo, student.Class, updatedStudent.Class)
		t.Fatalf("Updated student does not match the expected values")
	}


	// Test DeleteStudent
	err = repo.DeleteStudent(student.MongoID.Hex()) // Delete by ObjectID in hexadecimal format
	if err != nil {
		t.Fatalf("Failed to delete student: %v", err)
	}

	deletedStudent, err := repo.GetStudentByID(student.MongoID.Hex()) // Fetch by ObjectID in hexadecimal format
	if err != nil {
		t.Fatalf("Failed to get deleted student by ID: %v", err)
	}
	if deletedStudent != nil {
		t.Fatalf("Deleted student with ID %s still exists", student.MongoID.Hex())
	}
}
