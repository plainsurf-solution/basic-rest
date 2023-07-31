package repository

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"students/app/models"

	"github.com/go-redis/redis/v8"
)

// RedisStudentRepository represents a Redis implementation of the StudentRepository
type RedisStudentRepository struct {
	client *redis.Client
}

// NewRedisStudentRepository creates a new instance of RedisStudentRepository
func NewRedisStudentRepository(client *redis.Client) *RedisStudentRepository {
	return &RedisStudentRepository{
		client: client,
	}
}

// GetStudents retrieves all students from the Redis repository
func (r *RedisStudentRepository) GetStudents() ([]models.Student, error) {
	// Redis-specific implementation to fetch students
	// ...
	ctx := context.TODO()

	// Get the "students" key from Redis
	studentsJSON, err := r.client.Get(ctx, "students:").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// Key not found in Redis, return an empty slice of students
			return []models.Student{}, nil
		}
		return nil, err
	}

	// Unmarshal the JSON data back into a slice of students
	var students []models.Student
	err = json.Unmarshal([]byte(studentsJSON), &students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

// GetStudentByID retrieves a student by ID from the Redis repository
func (r *RedisStudentRepository) GetStudentByID(id string) (*models.Student, error) {
	// Redis-specific implementation to fetch a student by ID
	// ...
	ctx := context.TODO()

	// Get the student data by ID from Redis
	studentJSON, err := r.client.HGet(ctx, "student:"+id, "data").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// Key not found in Redis, return nil
			return nil, nil
		}
		return nil, err
	}

	// Unmarshal the JSON data back into a student object
	var student models.Student
	err = json.Unmarshal([]byte(studentJSON), &student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

// CreateStudent adds a new student to the Redis repository
func (r *RedisStudentRepository) CreateStudent(student *models.Student) error {
	// Redis-specific implementation to create a student
	// ...
	ctx := context.TODO()

	// Marshal the student data to JSON
	studentJSON, err := json.Marshal(student)
	if err != nil {
		return err
	}

	// Add the student data to Redis using HSET
	_, err = r.client.HSet(ctx, "student:"+strconv.Itoa(int(student.ID)), "data", studentJSON).Result()
	if err != nil {
		return err
	}

	return nil
}

// UpdateStudent updates a student in the Redis repository
func (r *RedisStudentRepository) UpdateStudent(student *models.Student) error {
	// Redis-specific implementation to update a student
	// ...
	ctx := context.TODO()

	// Marshal the updated student data to JSON
	studentJSON, err := json.Marshal(student)
	if err != nil {
		return err
	}

	// Update the student data in Redis using HSET
	_, err = r.client.HSet(ctx, "student:"+strconv.Itoa(int(student.ID)), "data", studentJSON).Result()
	if err != nil {
		return err
	}

	return nil
}

// DeleteStudent deletes a student by ID from the Redis repository
func (r *RedisStudentRepository) DeleteStudent(id string) error {
	// Redis-specific implementation to delete a student
	// ...
	ctx := context.TODO()

	// Delete the student data from Redis
	_, err := r.client.Del(ctx, "student:"+id).Result()
	if err != nil {
		return err
	}

	return nil
}
