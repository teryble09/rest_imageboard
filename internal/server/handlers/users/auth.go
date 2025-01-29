package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rest_imageboard/internal/storage"
	"rest_imageboard/internal/storage/query"

	"github.com/golang-jwt/jwt/v5"
)

// автоматически создает пользователя если такого нет, либо находит его в базе
// и возвращает вечный токен (что априори ужас)
// но пока пойдет
func GetTokenFromNamePassword(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := query.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		isInDB, err := query.UserIsInDB(db, &user)
		if err == storage.ErrPasswordDoesNotMatch {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !isInDB {
			err := query.CreateUser(db, &user)
			fmt.Println("Create: " + err.Error())
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Name
		tokenString, err := token.SignedString(storage.Secret)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(tokenString))
	}
}

// удаляет пользователя из базы, принимает токен
func DeleteAccount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		claims := token.Claims.(jwt.MapClaims)
		user := query.User{Name: claims["username"].(string)}
		err = query.DeleteUser(db, user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}