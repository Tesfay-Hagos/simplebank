package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tesfayprep/user_authentication/model"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := model.UserInfo{}
	json.NewDecoder(r.Body).Decode(&user)
	response := model.Register(user)
	json.NewEncoder(w).Encode(response)
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := model.UserInfo{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "invalid body")
		return
	}
	dbuser, isuser := model.GetUser(user.Email)
	if isuser {
		equal := bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(user.Password))
		if equal != nil {
			fmt.Fprintf(w, "can not authenticate this user")
			return
		} else {
			token, err := GenerateJWT(user.Username)
			if err != nil {
				fmt.Fprintf(w, "error in generating token")
			}

			fmt.Fprintf(w, token)
		}
	}
}

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	user := model.UserInfo{}
	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-Type", "application/json")
	message := model.ChngePassword(user.Email, user.Password)
	json.NewEncoder(w).Encode(message)
}
func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	err := ValidateToken(w, r)

	user := model.UserInfo{}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		users := model.GetAllUser()
		json.NewEncoder(w).Encode(users)
	}
}
func hashpassword(pw []byte) []byte {
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
