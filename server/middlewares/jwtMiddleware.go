package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

const jwtSecret1 = "your_secret_key"

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the "Authorization" header
		
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract the token from the "Authorization" header
		splitToken := strings.Split(tokenHeader, "Bearer ")
		if len(splitToken) != 2 {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		tokenString := splitToken[1]

		// Validate and parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method and return the secret key
			if token.Method.Alg() != jwt.SigningMethodHS256.Name {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler if the token is valid
		next(w, r)
	}
}
