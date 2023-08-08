package repository

import (
	"errors"
	"students/app/models"

	"gorm.io/gorm"
)

// MySQLStudentRepository represents a MySQL implementation of the StudentRepository
type MySQLStudentRepository struct {
	connection *gorm.DB
}

// NewMySQLStudentRepository creates a new instance of MySQLStudentRepository
func NewMySQLStudentRepository(connection *gorm.DB) *MySQLStudentRepository {
	return &MySQLStudentRepository{
		connection: connection,
	}
}

// GetStudents retrieves all students from the MySQL repository
func (r *MySQLStudentRepository) GetStudents() ([]models.Student, error) {
	var students []models.Student
	if err := r.connection.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// GetStudentByID retrieves a student by ID from the MySQL repository
func (r *MySQLStudentRepository) GetStudentByID(id string) (*models.Student, error) {
	var student models.Student
	if err := r.connection.Where("id = ?", id).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &student, nil
}

// CreateStudent adds a new student to the MySQL repository
func (r *MySQLStudentRepository) CreateStudent(student *models.Student) error {
	if student.RollNo == "" {
		return errors.New("rollno field is required")
	}
	if err := r.connection.Create(student).Error; err != nil {
		return err
	}
	return nil
}

// UpdateStudent updates a student in the MySQL repository
func (r *MySQLStudentRepository) UpdateStudent(student *models.Student) error {
	if err := r.connection.Model(student).Updates(student).Error; err != nil {
		return err
	}
	return nil
}

// DeleteStudent deletes a student by ID from the MySQL repository
func (r *MySQLStudentRepository) DeleteStudent(id string) error {
	if err := r.connection.Where("id = ?", id).Delete(&models.Student{}).Error; err != nil {
		return err
	}
	return nil
}
