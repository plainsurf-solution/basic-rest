package middlewares

import (
	"students/app/models"
	"time"

	"github.com/golang-jwt/jwt"
)

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
