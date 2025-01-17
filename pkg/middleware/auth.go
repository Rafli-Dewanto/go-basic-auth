// pkg/middleware/auth_middleware.go
package middleware

import (
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearerToken := r.Header.Get("Authorization")
			tokenString := bearerToken[7:] // get token string from "Bearer <token>"
			if tokenString == "" {
				utils.ErrorResponse(w, "Missing token", http.StatusUnauthorized)
				return
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				utils.ErrorResponse(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
