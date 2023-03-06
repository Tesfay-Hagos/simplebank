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

func TestUserregisterandlogin(t *testing.T) {
	server := controller.Newserver()
	t.Run("Test Registering users", func(t *testing.T) {
		Newuser := model.UserInfo{Username: "Bnet", Password: "Berut2121", Email: "bnetthagos@gmail.com", CreatedAt: time.Now()}
		//create request and response writer
		buff := convtobuff(Newuser)
		req := httptest.NewRequest(http.MethodPost, "/register", &buff)
		resp := httptest.NewRecorder()
		//initiate server
		server.Handler.ServeHTTP(resp, req)
		response := model.JsonResponse{}
		json.NewDecoder(resp.Body).Decode(&response)
		//check response
		if response.Type != "success" {
			t.Errorf("Test Failed")
		}
	})
	t.Run("Test login user and generate token", func(t *testing.T) {
		Newuser := model.UserInfo{Password: "Berut2121", Email: "bemnetthagos@gmail.com"}
		buff := convtobuff(Newuser)
		req := httptest.NewRequest(http.MethodPost, "/login", &buff)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		if resp.Body.String() == "" {
			t.Errorf("Test Failed")
		}
	})

}
func convtobuff(user model.UserInfo) bytes.Buffer {
	body, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	buff := bytes.NewBuffer(body)
	return *buff

}
