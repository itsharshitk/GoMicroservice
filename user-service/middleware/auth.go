package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("your-secret-key") // Same as in auth-service

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Optional: pass email in context if needed
		next.ServeHTTP(w, r)
	})
}
