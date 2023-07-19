// repository.go

package repository

import (
	"context"
	"errors"

	"students/app/models"
	"students/common"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		return NewMySQLStudentRepository(dbConfig.Connection), nil
	case "DatabaseTypePostgreSQL":
		// Initialize and return a PostgreSQL repository
		return NewPostgreSQLStudentRepository(dbConfig.Connection), nil
	default:
		return nil, errors.New("unsupported database type")
	}
}

// MongoDBStudentRepository represents a MongoDB implementation of the StudentRepository
type MongoDBStudentRepository struct {
	collection *mongo.Collection
}

// NewMongoDBStudentRepository creates a new instance of MongoDBStudentRepository
func NewMongoDBStudentRepository(client *mongo.Client) *MongoDBStudentRepository {
	return &MongoDBStudentRepository{
		collection: client.Database("students").Collection("collection"),
	}
}

// GetStudents retrieves all students from the MongoDB repository
func (r *MongoDBStudentRepository) GetStudents() ([]models.Student, error) {
	// MongoDB-specific implementation to fetch students
	// ...
	ctx := context.TODO()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []models.Student
	for cursor.Next(ctx) {
		var student models.Student
		err := cursor.Decode(&student)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

// GetStudentByID retrieves a student by ID from the MongoDB repository
func (r *MongoDBStudentRepository) GetStudentByID(id string) (*models.Student, error) {
	// MongoDB-specific implementation to fetch a student by ID
	// ...
	// Parse the given string ID into an ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	filter := bson.M{"_id": objID}

	var student models.Student
	err = r.collection.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &student, nil
}

// CreateStudent adds a new student to the MongoDB repository
func (r *MongoDBStudentRepository) CreateStudent(student *models.Student) error {
	// MongoDB-specific implementation to create a student
	// ...
	ctx := context.TODO()

	_, err := r.collection.InsertOne(ctx, student)
	if err != nil {
		return err
	}

	return nil
}

// UpdateStudent updates a student in the MongoDB repository
func (r *MongoDBStudentRepository) UpdateStudent(student *models.Student) error {
	// MongoDB-specific implementation to update a student
	// ...
	ctx := context.TODO()

	filter := bson.M{"id": student.ID}

	update := bson.M{
		"$set": bson.M{
			"name":   student.Name,
			"id":     student.ID,
			"roolno": student.RollNo,
			"class":  student.Class,
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteStudent deletes a student by ID from the MongoDB repository
func (r *MongoDBStudentRepository) DeleteStudent(id string) error {
	// MongoDB-specific implementation to delete a student
	// ...
	ctx := context.TODO()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	_, err = r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

// MySQLStudentRepository represents a MySQL implementation of the StudentRepository
type MySQLStudentRepository struct {
	connection string
}

// NewMySQLStudentRepository creates a new instance of MySQLStudentRepository
func NewMySQLStudentRepository(connection string) *MySQLStudentRepository {
	return &MySQLStudentRepository{
		connection: connection,
	}
}

// GetStudents retrieves all students from the MySQL repository
func (r *MySQLStudentRepository) GetStudents() ([]models.Student, error) {
	// MySQL-specific implementation to fetch students
	// ...
	return nil, nil
}

// GetStudentByID retrieves a student by ID from the MySQL repository
func (r *MySQLStudentRepository) GetStudentByID(id string) (*models.Student, error) {
	// MySQL-specific implementation to fetch a student by ID
	// ...
	return nil, nil
}

// CreateStudent adds a new student to the MySQL repository
func (r *MySQLStudentRepository) CreateStudent(student *models.Student) error {
	// MySQL-specific implementation to create a student
	// ...
	return nil
}

// UpdateStudent updates a student in the MySQL repository
func (r *MySQLStudentRepository) UpdateStudent(student *models.Student) error {
	// MySQL-specific implementation to update a student
	// ...
	return nil
}

// DeleteStudent deletes a student by ID from the MySQL repository
func (r *MySQLStudentRepository) DeleteStudent(id string) error {
	// MySQL-specific implementation to delete a student
	// ...
	return nil
}

// PostgreSQLStudentRepository represents a PostgreSQL implementation of the StudentRepository
type PostgreSQLStudentRepository struct {
	connection string
}

// NewPostgreSQLStudentRepository creates a new instance of PostgreSQLStudentRepository
func NewPostgreSQLStudentRepository(connection string) *PostgreSQLStudentRepository {
	return &PostgreSQLStudentRepository{
		connection: connection,
	}
}

// GetStudents retrieves all students from the PostgreSQL repository
func (r *PostgreSQLStudentRepository) GetStudents() ([]models.Student, error) {
	// PostgreSQL-specific implementation to fetch students
	// ...
	return nil, nil
}

// GetStudentByID retrieves a student by ID from the PostgreSQL repository
func (r *PostgreSQLStudentRepository) GetStudentByID(id string) (*models.Student, error) {
	// PostgreSQL-specific implementation to fetch a student by ID
	// ...
	return nil, nil
}

// CreateStudent adds a new student to the PostgreSQL repository
func (r *PostgreSQLStudentRepository) CreateStudent(student *models.Student) error {
	// PostgreSQL-specific implementation to create a student
	// ...
	return nil
}

// UpdateStudent updates a student in the PostgreSQL repository
func (r *PostgreSQLStudentRepository) UpdateStudent(student *models.Student) error {
	// PostgreSQL-specific implementation to update a student
	// ...
	return nil
}

// DeleteStudent deletes a student by ID from the PostgreSQL repository
func (r *PostgreSQLStudentRepository) DeleteStudent(id string) error {
	// PostgreSQL-specific implementation to delete a student
	// ...
	return nil
}

// // repository.go

// package repository

// import (
// 	"context"
// 	"errors"

// 	"students/app/models"
// 	"students/common"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // StudentRepository provides an interface to interact with the student data store
// type StudentRepository interface {
// 	GetStudents() ([]models.Student, error)
// 	GetStudentByID(id string) (*models.Student, error)
// 	CreateStudent(student *models.Student) error
// 	UpdateStudent(student *models.Student) error
// 	DeleteStudent(id string) error
// }

// // NewStudentRepository creates a new instance of the StudentRepository based on the database configuration
// func NewStudentRepository(dbConfig common.DatabaseConfig) (StudentRepository, error) {
// 	switch dbConfig.Type {
// 	case "DatabaseTypeMongoDB":
// 		mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConfig.Connection))
// 		if err != nil {
// 			return nil, err
// 		}
// 		return NewMongoDBStudentRepository(mongoClient), nil
// 	case "DatabaseTypeMySQL":
// 		// Initialize and return a MySQL repository
// 		return NewMySQLStudentRepository(dbConfig.Connection), nil
// 	case "DatabaseTypePostgreSQL":
// 		// Initialize and return a PostgreSQL repository
// 		return NewPostgreSQLStudentRepository(dbConfig.Connection), nil
// 	default:
// 		return nil, errors.New("unsupported database type")
// 	}
// }

// // MongoDBStudentRepository represents a MongoDB implementation of the StudentRepository
// type MongoDBStudentRepository struct {
// 	client *mongo.Client
// }

// // NewMongoDBStudentRepository creates a new instance of MongoDBStudentRepository
// func NewMongoDBStudentRepository(client *mongo.Client) *MongoDBStudentRepository {
// 	return &MongoDBStudentRepository{
// 		client: client,
// 	}
// }

// // GetStudents retrieves all students from the MongoDB repository
// func (r *MongoDBStudentRepository) GetStudents() ([]models.Student, error) {
// 	// MongoDB-specific implementation to fetch students
// 	// ...
// }

// // GetStudentByID retrieves a student by ID from the MongoDB repository
// func (r *MongoDBStudentRepository) GetStudentByID(id string) (*models.Student, error) {
// 	// MongoDB-specific implementation to fetch a student by ID
// 	// ...
// }

// // CreateStudent adds a new student to the MongoDB repository
// func (r *MongoDBStudentRepository) CreateStudent(student *models.Student) error {
// 	// MongoDB-specific implementation to create a student
// 	// ...
// }

// // UpdateStudent updates a student in the MongoDB repository
// func (r *MongoDBStudentRepository) UpdateStudent(student *models.Student) error {
// 	// MongoDB-specific implementation to update a student
// 	// ...
// }

// // DeleteStudent deletes a student by ID from the MongoDB repository
// func (r *MongoDBStudentRepository) DeleteStudent(id string) error {
// 	// MongoDB-specific implementation to delete a student
// 	// ...
// }

// // MySQLStudentRepository represents a MySQL implementation of the StudentRepository
// type MySQLStudentRepository struct {
// 	connection string
// }

// // NewMySQLStudentRepository creates a new instance of MySQLStudentRepository
// func NewMySQLStudentRepository(connection string) *MySQLStudentRepository {
// 	return &MySQLStudentRepository{
// 		connection: connection,
// 	}
// }

// // GetStudents retrieves all students from the MySQL repository
// func (r *MySQLStudentRepository) GetStudents() ([]models.Student, error) {
// 	// MySQL-specific implementation to fetch students
// 	// ...
// }

// // GetStudentByID retrieves a student by ID from the MySQL repository
// func (r *MySQLStudentRepository) GetStudentByID(id string) (*models.Student, error) {
// 	// MySQL-specific implementation to fetch a student by ID
// 	// ...
// }

// // CreateStudent adds a new student to the MySQL repository
// func (r *MySQLStudentRepository) CreateStudent(student *models.Student) error {
// 	// MySQL-specific implementation to create a student
// 	// ...
// }

// // UpdateStudent updates a student in the MySQL repository
// func (r *MySQLStudentRepository) UpdateStudent(student *models.Student) error {
// 	// MySQL-specific implementation to update a student
// 	// ...
// }

// // DeleteStudent deletes a student by ID from the MySQL repository
// func (r *MySQLStudentRepository) DeleteStudent(id string) error {
// 	// MySQL-specific implementation to delete a student
// 	// ...
// }

// // PostgreSQLStudentRepository represents a PostgreSQL implementation of the StudentRepository
// type PostgreSQLStudentRepository struct {
// 	connection string
// }

// // NewPostgreSQLStudentRepository creates a new instance of PostgreSQLStudentRepository
// func NewPostgreSQLStudentRepository(connection string) *PostgreSQLStudentRepository {
// 	return &PostgreSQLStudentRepository{
// 		connection: connection,
// 	}
// }

// // GetStudents retrieves all students from the PostgreSQL repository
// func (r *PostgreSQLStudentRepository) GetStudents() ([]models.Student, error) {
// 	// PostgreSQL-specific implementation to fetch students
// 	// ...
// }

// // GetStudentByID retrieves a student by ID from the PostgreSQL repository
// func (r *PostgreSQLStudentRepository) GetStudentByID(id string) (*models.Student, error) {
// 	// PostgreSQL-specific implementation to fetch a student by ID
// 	// ...
// }

// // CreateStudent adds a new student to the PostgreSQL repository
// func (r *PostgreSQLStudentRepository) CreateStudent(student *models.Student) error {
// 	// PostgreSQL-specific implementation to create a student
// 	// ...
// }

// // UpdateStudent updates a student in the PostgreSQL repository
// func (r *PostgreSQLStudentRepository) UpdateStudent(student *models.Student) error {
// 	// PostgreSQL-specific implementation to update a student
// 	// ...
// }

// // DeleteStudent deletes a student by ID from the PostgreSQL repository
// func (r *PostgreSQLStudentRepository) DeleteStudent(id string) error {
// 	// PostgreSQL-specific implementation to delete a student
// 	// ...
// }
