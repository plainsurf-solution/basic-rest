// services.go

package services

import (
	"students/app/models"
	"students/pkg/repository"
)

// StudentService handles student-related operations
type StudentService struct {
	studentRepo repository.StudentRepository
}

// NewStudentService creates a new instance of StudentService
func NewStudentService(studentRepo repository.StudentRepository) *StudentService {
	return &StudentService{
		studentRepo: studentRepo,
	}
}

// GetStudents retrieves all students
func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.studentRepo.GetStudents()
}

// GetStudentByID retrieves a student by ID
func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.studentRepo.GetStudentByID(id)
}

// CreateStudent creates a new student
func (s *StudentService) CreateStudent(student *models.Student) error {
	return s.studentRepo.CreateStudent(student)
}

// UpdateStudent updates a student
func (s *StudentService) UpdateStudent(student *models.Student) error {
	return s.studentRepo.UpdateStudent(student)
}

// DeleteStudent deletes a student by ID
func (s *StudentService) DeleteStudent(id string) error {
	return s.studentRepo.DeleteStudent(id)
}
