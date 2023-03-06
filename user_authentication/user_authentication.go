package main

import (
	"net/http"
	"tesfayprep/user_authentication/controller"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/register", controller.RegisterHandler)
	router.Handle("/login", controller.LoginMiddleware(http.HandlerFunc(controller.LoginHandler)))
	router.HandleFunc("/getalluser", controller.GetAllUserHandler)
	http.ListenAndServe(":8080", router)
}
