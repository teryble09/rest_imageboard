package middleware

import (
	"net/http"
	"rest_imageboard/internal/storage"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
            	http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            	return
       	 	}
			// Bearer 
			tokenString = tokenString[7:]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
			return storage.Secret, nil
			})

			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		},
	)
}