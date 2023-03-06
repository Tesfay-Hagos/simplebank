package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"tesfayprep/user_authentication/controller"
	"tesfayprep/user_authentication/model"
	"testing"
	"time"
)

func TestRegiste(t *testing.T) {
	Newuser := model.UserInfo{Username: "Bemnet", Password: "Berut2121", Email: "bemnetthagos@gmail.com", CreatedAt: time.Now()}
	buff := convtobuff(Newuser)
	router := http.NewServeMux()
	router.HandleFunc("/register", controller.RegisterHandler)
	req := httptest.NewRequest(http.MethodPost, "/register", &buff)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	response := model.JsonResponse{}
	json.NewDecoder(resp.Body).Decode(&response)
	if response.Type != "success" {
		t.Errorf("Test Failed")
	}

}
func TestLogin(t *testing.T) {
	Newuser := model.UserInfo{Password: "Berut2121", Email: "bemnetthagos@gmail.com"}
	buff := convtobuff(Newuser)
	router := http.NewServeMux()
	router.Handle("/login", controller.LoginMiddleware(http.HandlerFunc(controller.LoginHandler)))
	req := httptest.NewRequest(http.MethodPost, "/login", &buff)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Body.String() == "" {
		t.Errorf("Test Failed")
	}

}
func convtobuff(user model.UserInfo) bytes.Buffer {
	body, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	buff := bytes.NewBuffer(body)
	return *buff

}
