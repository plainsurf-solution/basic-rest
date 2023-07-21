// controllers.go

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"students/app/models"
	"students/app/services"

	"github.com/gorilla/mux"
)

// StudentController handles student-related HTTP requests
type StudentController struct {
	studentService *services.StudentService
}

// NewStudentController creates a new instance of StudentController
func NewStudentController(studentService *services.StudentService) *StudentController {
	return &StudentController{
		studentService: studentService,
	}
}

// GetStudents handles the HTTP GET request to retrieve all students
func (c *StudentController) GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := c.studentService.GetStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(students)
}

// GetStudentByID handles the HTTP GET request to retrieve a student by ID
func (c *StudentController) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	// studentID := r.URL.Query().Get("id")
	params := mux.Vars(r)
	studentID := params["id"]

	student, err := c.studentService.GetStudentByID(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

// CreateStudent handles the HTTP POST request to create a new student
func (c *StudentController) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		fmt.Println("here", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.studentService.CreateStudent(&student)
	if err != nil {
		fmt.Println("5454", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// UpdateStudent handles the HTTP PUT request to update an existing student
func (c *StudentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.studentService.UpdateStudent(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(student)
}

// DeleteStudent handles the HTTP DELETE request to delete a student by ID
func (c *StudentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	studentID := params["id"]

	err := c.studentService.DeleteStudent(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
