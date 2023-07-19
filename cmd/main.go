// main.go

package main

import (
	"log"
	"net/http"

	"students/app/controllers"
	"students/app/services"
	"students/common"
	"students/pkg/repository"
	"students/server/router"

	"github.com/gorilla/mux"
)

func main() {

	dbConfig, err := common.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	// Create instances of the repository, service, and controller
	studentRepo, err := repository.NewStudentRepository(*dbConfig) // Replace with your repository implementation
	studentService := services.NewStudentService(studentRepo)
	studentController := controllers.NewStudentController(studentService)

	// Create a new router
	r := mux.NewRouter()

	// Register routes
	router.RegisterRoutes(r, studentController)

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
