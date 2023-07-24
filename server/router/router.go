// router.go

package router

import (
	"net/http"
	"students/app/controllers"
	"students/server/middlewares"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers the routes and their corresponding handlers
func RegisterRoutes(router *mux.Router, studentController *controllers.StudentController) {
	// Use the AuthMiddleware for "/students" with GET, PUT, and DELETE methods.
	router.Path("/students").Methods(http.MethodGet).Handler(middlewares.JwtMiddleware(http.HandlerFunc(studentController.GetStudents)))
	router.Path("/students").Methods(http.MethodPut).Handler(middlewares.JwtMiddleware(http.HandlerFunc(studentController.UpdateStudent)))
	router.Path("/students/{id}").Methods(http.MethodDelete).Handler(middlewares.JwtMiddleware(http.HandlerFunc(studentController.DeleteStudent)))

	// Use the AuthMiddleware for "/students/{id}" with GET and DELETE methods.
	router.Path("/students/{id}").Methods(http.MethodGet).Handler(middlewares.JwtMiddleware(http.HandlerFunc(studentController.GetStudentByID)))

	// The "create" route remains unprotected.
	router.HandleFunc("/students", studentController.CreateStudent).Methods(http.MethodPost)
	
}



