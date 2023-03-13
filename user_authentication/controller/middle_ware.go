package controller

import (
	"fmt"
	"net/http"
	"tesfayprep/user_authentication/model"
)

func LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			next.ServeHTTP(w, r)
		case "GET":
			fmt.Fprintf(w, "only POST methods is allowed.")
			return
		}
	})
}

func ChangePasswordHandlermiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ValidateToken(w, r)
		model.CheckErr(err)
		if err == nil {
			next.ServeHTTP(w, r)

		}

	})
}
func GetAllUSerHandlermiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ValidateToken(w, r)
		if err != nil {
			fmt.Fprintf(w, "Token invalid")
		} else {
			next.ServeHTTP(w, r)

		}
	})
}
func ResetpasswordRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func ResetpasswordMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
