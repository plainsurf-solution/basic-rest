// router.go

package router

import (
	"net/http"
	"students/app/controllers"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers the routes and their corresponding handlers
func RegisterRoutes(router *mux.Router, studentController *controllers.StudentController) {
	router.HandleFunc("/students", studentController.GetStudents).Methods(http.MethodGet)
	router.HandleFunc("/students/{id}", studentController.GetStudentByID).Methods(http.MethodGet)
	router.HandleFunc("/students", studentController.CreateStudent).Methods(http.MethodPost)
	router.HandleFunc("/students", studentController.UpdateStudent).Methods(http.MethodPut)
	router.HandleFunc("/students/{id}", studentController.DeleteStudent).Methods(http.MethodDelete)
}
