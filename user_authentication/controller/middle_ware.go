package controller

import (
	"fmt"
	"net/http"
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
