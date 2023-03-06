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
		next.ServeHTTP(w, r)
	})
}
