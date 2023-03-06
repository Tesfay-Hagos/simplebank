package controller

import "net/http"

type Server struct {
	http.Server
}

func Newserver() *Server {
	serv := new(Server)
	router := http.NewServeMux()
	router.HandleFunc("/register", RegisterHandler)
	router.Handle("/login", LoginMiddleware(http.HandlerFunc(LoginHandler)))
	router.Handle("/changepassword", ChangePasswordHandlermiddleware(http.HandlerFunc(ChangePasswordHandler)))
	router.HandleFunc("/getalluser", GetAllUserHandler)
	serv.Handler = router
	return serv
}
