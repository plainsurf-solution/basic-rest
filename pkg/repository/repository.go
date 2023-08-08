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

	"gorm.io/driver/mysql"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
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
		  // Initialize and return a MySQL repository
		  MySQLClient, err := gorm.Open(mysql.Open(dbConfig.Connection), &gorm.Config{})
		  if err != nil {
			  panic("Failed to connect to DB")
		  }
	  
		  fmt.Println("Success! MySQL Database connected")
	  
		  return NewMySQLStudentRepository(MySQLClient), nil
	

	case "DatabaseTypePostgreSQL":
		// Initialize and return a PostgreSQL repository
		PostgreSQLClient, err := gorm.Open(postgres.Open(dbConfig.Connection), &gorm.Config{})
		if err != nil {
			panic("Falied to connect to DB")
		}
		PostgreSQLClient.AutoMigrate(&models.Student{}) // This will sync tables of database and struct and matches the feilds
		return NewPostgreSQLStudentRepository(PostgreSQLClient), nil
	case "DatabaseTypeRedis":
		// Initialize and return a Redis repository
		redisClient := redis.NewClient(&redis.Options{
			Addr:     dbConfig.Connection,
			Password: "", // If your Redis instance has authentication enabled, provide the password here
			DB:       0,  // Use default DB
		})
	
		// Check the connection
		_, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			return nil, err
		}
	
		return NewRedisStudentRepository(redisClient), nil
	default:
		return nil, errors.New("unsupported database type")
	}
}

