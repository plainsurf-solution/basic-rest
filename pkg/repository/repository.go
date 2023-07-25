// repository.go

package repository

import (
	"context"
	// "encoding/json"
	"errors"
	"fmt"

	"students/app/models"
	"students/common"

	// "golang.org/x/crypto/bcrypt"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
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

// // MongoDBStudentRepository represents a MongoDB implementation of the StudentRepository
// type MongoDBStudentRepository struct {
// 	collection *mongo.Collection
// }

// // NewMongoDBStudentRepository creates a new instance of MongoDBStudentRepository
// func NewMongoDBStudentRepository(client *mongo.Client) *MongoDBStudentRepository {
// 	return &MongoDBStudentRepository{
// 		collection: client.Database("students").Collection("collection"),
// 	}
// }

// // GetStudents retrieves all students from the MongoDB repository
// func (r *MongoDBStudentRepository) GetStudents() ([]models.Student, error) {
// 	// MongoDB-specific implementation to fetch students
// 	// ...
// 	ctx := context.TODO()

// 	cursor, err := r.collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var students []models.Student
// 	for cursor.Next(ctx) {
// 		var student models.Student
// 		err := cursor.Decode(&student)
// 		if err != nil {
// 			return nil, err
// 		}

// 		students = append(students, student)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return students, nil
// }

// // GetStudentByID retrieves a student by ID from the MongoDB repository
// func (r *MongoDBStudentRepository) GetStudentByID(id string) (*models.Student, error) {
// 	// MongoDB-specific implementation to fetch a student by ID
// 	// ...
// 	// Parse the given string ID into an ObjectID
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ctx := context.TODO()

// 	filter := bson.M{"_id": objID}

// 	var student models.Student
// 	err = r.collection.FindOne(ctx, filter).Decode(&student)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &student, nil
// }

// // CreateStudent adds a new student to the MongoDB repository
// func (r *MongoDBStudentRepository) CreateStudent(student *models.Student) error {
// 	// MongoDB-specific implementation to create a student
// 	// ...
// 	ctx := context.TODO()

// 	// Hash the password using bcrypt
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	// Replace the plain-text password with the hashed password
// 	student.Password = string(hashedPassword)

// 	// Store the updated student object in the MongoDB collection
// 	_, err = r.collection.InsertOne(ctx, student)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // UpdateStudent updates a student in the MongoDB repository
// func (r *MongoDBStudentRepository) UpdateStudent(student *models.Student) error {
// 	// MongoDB-specific implementation to update a student
// 	// ...
// 	ctx := context.TODO()

// 	filter := bson.M{"id": student.ID}

// 	update := bson.M{
// 		"$set": bson.M{
// 			"name":   student.Name,
// 			"id":     student.ID,
// 			"roolno": student.RollNo,
// 			"class":  student.Class,
// 		},
// 	}

// 	_, err := r.collection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // DeleteStudent deletes a student by ID from the MongoDB repository
// func (r *MongoDBStudentRepository) DeleteStudent(id string) error {
// 	// MongoDB-specific implementation to delete a student
// 	// ...
// 	ctx := context.TODO()
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	filter := bson.M{"_id": objID}

// 	_, err = r.collection.DeleteOne(ctx, filter)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// MySQLStudentRepository represents a MySQL implementation of the StudentRepository
// type MySQLStudentRepository struct {
// 	connection *sql.DB
// }

// // NewMySQLStudentRepository creates a new instance of MySQLStudentRepository
// func NewMySQLStudentRepository(connection *sql.DB) *MySQLStudentRepository {
// 	return &MySQLStudentRepository{
// 		connection: connection,
// 	}
// }

// // GetStudents retrieves all students from the MySQL repository
// func (r *MySQLStudentRepository) GetStudents() ([]models.Student, error) {
// 	var students []models.Student
// 	query := "SELECT * FROM students"
// 	rows, err := r.connection.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var student models.Student
// 		var optionalSubjectsJSON []byte
// 		if err := rows.Scan(&student.ID, &student.Name, &student.RollNo, &student.Class, &optionalSubjectsJSON, &student.StudentRank); err != nil {
// 			return nil, err
// 		}

// 		if err := json.Unmarshal(optionalSubjectsJSON, &student.OptionalSubjects); err != nil {
// 			return nil, err
// 		}

// 		students = append(students, student)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return students, nil
// }

// // GetStudentByID retrieves a student by ID from the MySQL repository
// func (r *MySQLStudentRepository) GetStudentByID(id string) (*models.Student, error) {
// 	var student models.Student
// 	query := "SELECT * FROM students WHERE id = ?"
// 	row := r.connection.QueryRow(query, id)
// 	var optionalSubjectsJSON []byte
// 	if err := row.Scan(&student.ID, &student.Name, &student.RollNo, &student.Class, &optionalSubjectsJSON, &student.StudentRank); err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	if err := json.Unmarshal(optionalSubjectsJSON, &student.OptionalSubjects); err != nil {
// 		return nil, err
// 	}

// 	return &student, nil
// }

// // CreateStudent adds a new student to the MySQL repository
// func (r *MySQLStudentRepository) CreateStudent(student *models.Student) error {
// 	query := "INSERT INTO students (name, rollno, class, optional_subjects, student_rank) VALUES (?, ?, ?, ?, ?)"
// 	optionalSubjectsJSON, err := json.Marshal(student.OptionalSubjects)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.connection.Exec(query, student.Name, student.RollNo, student.Class, optionalSubjectsJSON, student.StudentRank)
// 	if err != nil {
// 		return err
// 	}

// 	studentID, err := result.LastInsertId()
// 	if err != nil {
// 		return err
// 	}

// 	student.ID = uint(studentID)

// 	return nil
// }

// // UpdateStudent updates a student in the MySQL repository
// func (r *MySQLStudentRepository) UpdateStudent(student *models.Student) error {
// 	query := "UPDATE students SET name = ?, rollno = ?, class = ?, optional_subjects = ?, student_rank = ? WHERE id = ?"

// 	optionalSubjectsJSON, err := json.Marshal(student.OptionalSubjects)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = r.connection.Exec(query, student.Name, student.RollNo, student.Class, optionalSubjectsJSON, student.StudentRank, student.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // DeleteStudent deletes a student by ID from the MySQL repository
// func (r *MySQLStudentRepository) DeleteStudent(id string) error {
// 	query := "DELETE FROM students WHERE id = ?"
// 	_, err := r.connection.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// PostgreSQLStudentRepository represents a PostgreSQL implementation of the StudentRepository
// type PostgreSQLStudentRepository struct {
// 	connection *gorm.DB
// }

// // NewPostgreSQLStudentRepository creates a new instance of PostgreSQLStudentRepository
// func NewPostgreSQLStudentRepository(connection *gorm.DB) *PostgreSQLStudentRepository {
// 	return &PostgreSQLStudentRepository{
// 		connection: connection,
// 	}
// }

// // GetStudents retrieves all students from the PostgreSQL repository
// func (r *PostgreSQLStudentRepository) GetStudents() ([]models.Student, error) {
// 	var students []models.Student
// 	if err := r.connection.Find(&students).Error; err != nil {
// 		return nil, err
// 	}
// 	return students, nil
// }

// // GetStudentByID retrieves a student by ID from the PostgreSQL repository
// func (r *PostgreSQLStudentRepository) GetStudentByID(id string) (*models.Student, error) {
// 	var student models.Student
// 	if err := r.connection.Where("id = ?", id).First(&student).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &student, nil
// }

// // CreateStudent adds a new student to the PostgreSQL repository
// func (r *PostgreSQLStudentRepository) CreateStudent(student *models.Student) error {
// 	fmt.Println("students", student)
// 	if err := r.connection.Create(student).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// // UpdateStudent updates a student in the PostgreSQL repository
// func (r *PostgreSQLStudentRepository) UpdateStudent(student *models.Student) error {
// 	if err := r.connection.Model(student).Updates(student).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// // DeleteStudent deletes a student by ID from the PostgreSQL repository
// func (r *PostgreSQLStudentRepository) DeleteStudent(id string) error {
// 	if err := r.connection.Where("id = ?", id).Delete(&models.Student{}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
