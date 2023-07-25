package repository

import (
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "encoding/json"
	"students/app/models"
	// "students/common"

	"errors"
	"fmt"
)

// PostgreSQLStudentRepository represents a PostgreSQL implementation of the StudentRepository
type PostgreSQLStudentRepository struct {
	connection *gorm.DB
}

// NewPostgreSQLStudentRepository creates a new instance of PostgreSQLStudentRepository
func NewPostgreSQLStudentRepository(connection *gorm.DB) *PostgreSQLStudentRepository {
	return &PostgreSQLStudentRepository{
		connection: connection,
	}
}

// GetStudents retrieves all students from the PostgreSQL repository
func (r *PostgreSQLStudentRepository) GetStudents() ([]models.Student, error) {
	var students []models.Student
	if err := r.connection.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// GetStudentByID retrieves a student by ID from the PostgreSQL repository
func (r *PostgreSQLStudentRepository) GetStudentByID(id string) (*models.Student, error) {
	var student models.Student
	if err := r.connection.Where("id = ?", id).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &student, nil
}

// CreateStudent adds a new student to the PostgreSQL repository
func (r *PostgreSQLStudentRepository) CreateStudent(student *models.Student) error {
	fmt.Println("students", student)
	if err := r.connection.Create(student).Error; err != nil {
		return err
	}
	return nil
}

// UpdateStudent updates a student in the PostgreSQL repository
func (r *PostgreSQLStudentRepository) UpdateStudent(student *models.Student) error {
	if err := r.connection.Model(student).Updates(student).Error; err != nil {
		return err
	}
	return nil
}

// DeleteStudent deletes a student by ID from the PostgreSQL repository
func (r *PostgreSQLStudentRepository) DeleteStudent(id string) error {
	if err := r.connection.Where("id = ?", id).Delete(&models.Student{}).Error; err != nil {
		return err
	}
	return nil
}