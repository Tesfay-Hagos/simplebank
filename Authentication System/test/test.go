package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Tesfay-Hagos/model"
)

func TestRegisterUser(t *testing.T) {
	Newuser := model.User{Username: "TH", Password: "123",
		Lastname: "Tesfay", Firstname: "Hagos", Email: "tesfya.hagos1421@gmail.com"}
	requestBody, err := json.Marshal(Newuser)
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
	respo := httptest.NewRecorder()
	model.RegisterUser(respo, req)

	got := model.GetUser(Newuser.Username, Newuser.Password)
	asserpassword(t, got.Password, Newuser.Password)
}
func asserpassword(t testing.TB, got, want string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%s, want:%s", got, want)
	}
}
