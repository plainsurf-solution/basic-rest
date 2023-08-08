package repository_test

import (
	"testing"

	"students/app/models"
	"students/pkg/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDBMYSQL() (*gorm.DB, error) {
	// Replace the connection details with your test MySQL database configuration
	dsn := "root:@123jkl@tcp(localhost:3306)/studentdbtest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Student{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestMySQLStudentRepository(t *testing.T) {
	db, err := setupTestDBMYSQL()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}

	// Create a new MySQLStudentRepository instance with the test database connection
	repo := repository.NewMySQLStudentRepository(db)

	// Test CreateStudent
	student := &models.Student{
		ID:          1,
		Email:       "test@example.com",
		Password:    "password123",
		Name:        "John Doe",
		RollNo:      "12345",
		Class:       "10A",
		StudentRank: 1,
	}
	err = repo.CreateStudent(student)
	if err != nil {
		t.Fatalf("Failed to create student: %v", err)
	}

	// Test GetStudents
	students, err := repo.GetStudents()
	if err != nil {
		t.Fatalf("Failed to get students: %v", err)
	}
	if len(students) != 1 {
		t.Fatalf("Expected 1 student, got %d", len(students))
	}
	if students[0].Email != "test@example.com" {
		t.Fatalf("Expected email to be 'test@example.com', got '%s'", students[0].Email)
	}

	// Test GetStudentByID
	foundStudent, err := repo.GetStudentByID("1")
	if err != nil {
		t.Fatalf("Failed to get student by ID: %v", err)
	}
	if foundStudent == nil {
		t.Fatalf("Expected a non-nil student, got nil")
	}
	if foundStudent.Email != "test@example.com" {
		t.Fatalf("Expected email to be 'test@example.com', got '%s'", foundStudent.Email)
	}

	// Test UpdateStudent
	student.Email = "updated@example.com"
	err = repo.UpdateStudent(student)
	if err != nil {
		t.Fatalf("Failed to update student: %v", err)
	}

	// Test GetStudentByID after updating
	updatedStudent, err := repo.GetStudentByID("1")
	if err != nil {
		t.Fatalf("Failed to get student by ID after updating: %v", err)
	}
	if updatedStudent == nil {
		t.Fatalf("Expected a non-nil student after updating, got nil")
	}
	if updatedStudent.Email != "updated@example.com" {
		t.Fatalf("Expected email to be 'updated@example.com', got '%s'", updatedStudent.Email)
	}

	// Test DeleteStudent
	err = repo.DeleteStudent("1")
	if err != nil {
		t.Fatalf("Failed to delete student: %v", err)
	}

	// Test GetStudentByID after deleting
	deletedStudent, err := repo.GetStudentByID("1")
	if err != nil {
		t.Fatalf("Failed to get student by ID after deleting: %v", err)
	}
	if deletedStudent != nil {
		t.Fatalf("Expected a nil student after deleting, got a non-nil student")
	}
}
