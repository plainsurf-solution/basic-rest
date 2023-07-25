package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"students/app/models"
)

type MySQLStudentRepository struct {
	connection *sql.DB
}

// NewMySQLStudentRepository creates a new instance of MySQLStudentRepository
func NewMySQLStudentRepository(connection *sql.DB) *MySQLStudentRepository {
	return &MySQLStudentRepository{
		connection: connection,
	}
}

// GetStudents retrieves all students from the MySQL repository
func (r *MySQLStudentRepository) GetStudents() ([]models.Student, error) {
	var students []models.Student
	query := "SELECT * FROM students"
	rows, err := r.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		var optionalSubjectsJSON []byte
		if err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Password, &student.RollNo, &student.Class, &optionalSubjectsJSON, &student.StudentRank); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(optionalSubjectsJSON, &student.OptionalSubjects); err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}


// GetStudentByID retrieves a student by ID from the MySQL repository
func (r *MySQLStudentRepository) GetStudentByID(id string) (*models.Student, error) {
	var student models.Student
	query := "SELECT * FROM students WHERE id = ?"
	row := r.connection.QueryRow(query, id)
	var optionalSubjectsJSON []byte
	if err := row.Scan(&student.ID, &student.Name, &student.Email, &student.Password, &student.RollNo, &student.Class, &optionalSubjectsJSON, &student.StudentRank); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(optionalSubjectsJSON, &student.OptionalSubjects); err != nil {
		return nil, err
	}

	return &student, nil
}


// CreateStudent adds a new student to the MySQL repository
// CreateStudent adds a new student to the MySQL repository
func (r *MySQLStudentRepository) CreateStudent(student *models.Student) error {
	query := "INSERT INTO students (name, email, password, rollno, class, optional_subjects, student_rank) VALUES (?, ?, ?, ?, ?, ?, ?)"
	optionalSubjectsJSON, err := json.Marshal(student.OptionalSubjects)
	if err != nil {
		return err
	}

	result, err := r.connection.Exec(query, student.Name, student.Email, student.Password, student.RollNo, student.Class, optionalSubjectsJSON, student.StudentRank)
	if err != nil {
		return err
	}

	studentID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	student.ID = uint(studentID)

	return nil
}


// UpdateStudent updates a student in the MySQL repository
func (r *MySQLStudentRepository) UpdateStudent(student *models.Student) error {
	query := "UPDATE students SET name = ?, email = ?, password = ?, rollno = ?, class = ?, optional_subjects = ?, student_rank = ? WHERE id = ?"

	optionalSubjectsJSON, err := json.Marshal(student.OptionalSubjects)
	if err != nil {
		return err
	}

	_, err = r.connection.Exec(query, student.Name, student.Email, student.Password, student.RollNo, student.Class, optionalSubjectsJSON, student.StudentRank, student.ID)
	if err != nil {
		return err
	}

	return nil
}


// DeleteStudent deletes a student by ID from the MySQL repository
func (r *MySQLStudentRepository) DeleteStudent(id string) error {
	query := "DELETE FROM students WHERE id = ?"
	_, err := r.connection.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}