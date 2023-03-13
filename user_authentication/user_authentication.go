package main

import (
	"net/http"
	"tesfayprep/user_authentication/controller"
)

func main() {
	server := controller.Newserver()
	http.ListenAndServe(":8080", server.Handler)
}
