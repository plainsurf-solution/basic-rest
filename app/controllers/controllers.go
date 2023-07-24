// controllers.go

package controllers

import (
	"encoding/json"
	"net/http"

	"students/app/models"
	"students/app/services"

	"github.com/gorilla/mux"

	"time"

	"github.com/golang-jwt/jwt"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.studentService.CreateStudent(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate and sign the JWT token
	token, err := generateJWTToken(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the token in the response header
	w.Header().Set("Authorization", "Bearer "+token)

	// Respond with the student data (if needed)
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




const jwtSecret = "your_secret_key"

func generateJWTToken(student *models.Student) (string, error) {
	// Set the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours

	// Create the claims for the token
	claims := jwt.MapClaims{
		"id":    student.ID,
		"email": student.Email,
		"name":  student.Name,
		// Add other claims here as needed
		"exp": expirationTime.Unix(),
	}

	// Create the token using the claims and the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}