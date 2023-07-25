// repository.go

package repository

import (
	"context"
	"errors"
	"fmt"

	"students/app/models"
	"students/common"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// StudentRepository provides an interface to interact with the student data store
type StudentRepository interface {
	GetStudents() ([]models.Student, error)
	GetStudentByID(id string) (*models.Student, error)
	CreateStudent(student *models.Student) error
	UpdateStudent(student *models.Student) error
	DeleteStudent(id string) error
}

// NewStudentRepository creates a new instance of the StudentRepository based on the database configuration
func NewStudentRepository(dbConfig common.DatabaseConfig) (StudentRepository, error) {
	switch dbConfig.Type {

	case "DatabaseTypeMongoDB":
		mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConfig.Connection))
		if err != nil {
			return nil, err
		}
		return NewMongoDBStudentRepository(mongoClient), nil

	case "DatabaseTypeMySQL":
		// Initialize and return a MySQL repository
		MySQLClient, err := sql.Open("mysql", dbConfig.Connection)
		if err != nil {
			panic(err.Error())
		}
		// defer MySQLClient.Close()
		fmt.Println("Success! mySQL Database connected")

		// Manually create the 'students' table if it doesn't exist
		createTableQuery := `
		CREATE TABLE IF NOT EXISTS students (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			rollno VARCHAR(20) NOT NULL,
			class VARCHAR(50) NOT NULL,
			optional_subjects JSON,
			student_rank INT
		  );
		  
`

		_, err = MySQLClient.Exec(createTableQuery)
		if err != nil {
			fmt.Println(err,"err")
			panic(err.Error())
		}

		return NewMySQLStudentRepository(MySQLClient), nil

	case "DatabaseTypePostgreSQL":
		// Initialize and return a PostgreSQL repository
		PostgreSQLClient, err := gorm.Open(postgres.Open(dbConfig.Connection), &gorm.Config{})
		if err != nil {
			panic("Falied to connect to DB")
		}
		PostgreSQLClient.AutoMigrate(&models.Student{}) // This will sync tables of database and struct and matches the feilds
		return NewPostgreSQLStudentRepository(PostgreSQLClient), nil

	default:
		return nil, errors.New("unsupported database type")
	}
}

