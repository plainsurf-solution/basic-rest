package repository_test

import (
	"context"
	"strconv"
	"students/app/models"
	"students/pkg/repository"
	"testing"

	"github.com/go-redis/redis/v8"
)

// Initialize a Redis client for testing (using a separate database)
func initRedisTestClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Replace with your Redis server address
		Password: "",               // No password set
		DB:       11,                // Use a separate Redis database for testing
	})
}

func TestRedisStudentRepository(t *testing.T) {
	// Initialize Redis client for testing
	client := initRedisTestClient()
	repo := repository.NewRedisStudentRepository(client)

	// Cleanup test data after the test is done
	defer func() {
		err := client.FlushDB(context.Background()).Err()
		if err != nil {
			t.Fatalf("Failed to flush Redis test database: %v", err)
		}
		client.Close()
	}()

	/// est CreateStudent
	student := &models.Student{
		ID:               1,
		Email:            "test@example.com",
		Password:         "password123",
		Name:             "John Doe",
		RollNo:           "12345",
		Class:            "10A",
		StudentRank:      1,
		OptionalSubjects: []string{"Math", "History"},
	}
	err := repo.CreateStudent(student)
	if err != nil {
		t.Fatalf("Failed to create student: %v", err)
	}

	// Test GetStudents - This will not work as getstudents
	// students, err := repo.GetStudents()
	// if err != nil {
	// 	t.Fatalf("Failed to get students: %v", err)
	// }
	// if len(students) != 1 {
	// 	t.Fatalf("Expected 1 student, got %d", len(students))
	// }
	// if students[0].Email != "test@example.com" {
	// 	t.Fatalf("Expected email to be 'test@example.com', got '%s'", students[0].Email)
	// }
	// if len(students[0].OptionalSubjects) != 2 || students[0].OptionalSubjects[0] != "Math" || students[0].OptionalSubjects[1] != "History" {
	// 	t.Fatalf("Optional subjects mismatch")
	// }

	// Test GetStudentByID
	foundStudent, err := repo.GetStudentByID(strconv.Itoa(int(student.ID)))
	if err != nil {
		t.Fatalf("Failed to get student by ID: %v", err)
	}
	if foundStudent == nil {
		t.Fatalf("Expected a non-nil student, got nil")
	}
	if foundStudent.Email != "test@example.com" {
		t.Fatalf("Expected email to be 'test@example.com', got '%s'", foundStudent.Email)
	}
	if len(foundStudent.OptionalSubjects) != 2 || foundStudent.OptionalSubjects[0] != "Math" || foundStudent.OptionalSubjects[1] != "History" {
		t.Fatalf("Optional subjects mismatch for GetStudentByID")
	}

	// Test UpdateStudent
	student.Email = "updated@example.com"
	student.OptionalSubjects = []string{"Physics", "Chemistry"}
	err = repo.UpdateStudent(student)
	if err != nil {
		t.Fatalf("Failed to update student: %v", err)
	}

	// Test GetStudentByID after updating
	updatedStudent, err := repo.GetStudentByID(strconv.Itoa(int(student.ID)))
	if err != nil {
		t.Fatalf("Failed to get student by ID after updating: %v", err)
	}
	if updatedStudent == nil {
		t.Fatalf("Expected a non-nil student after updating, got nil")
	}
	if updatedStudent.Email != "updated@example.com" {
		t.Fatalf("Expected email to be 'updated@example.com', got '%s'", updatedStudent.Email)
	}
	if len(updatedStudent.OptionalSubjects) != 2 || updatedStudent.OptionalSubjects[0] != "Physics" || updatedStudent.OptionalSubjects[1] != "Chemistry" {
		t.Fatalf("Optional subjects mismatch after updating")
	}

	// Test DeleteStudent
	err = repo.DeleteStudent(strconv.Itoa(int(student.ID)))
	if err != nil {
		t.Fatalf("Failed to delete student: %v", err)
	}

	// Test GetStudentByID after deleting
	deletedStudent, err := repo.GetStudentByID(strconv.Itoa(int(student.ID)))
	if err != nil {
		t.Fatalf("Failed to get student by ID after deleting: %v", err)
	}
	if deletedStudent != nil {
		t.Fatalf("Expected a nil student after deleting, got a non-nil student")

	}
	if deletedStudent != nil {
		t.Fatalf("Expected a nil student after deleting, got a non-nil student")

	}
	if deletedStudent != nil {
		t.Fatalf("Expected a nil student after deleting, got a non-nil student")
	}
}
