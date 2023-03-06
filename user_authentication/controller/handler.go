package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tesfayprep/user_authentication/model"
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
		if dbuser.Password == "" || dbuser.Password != user.Password {
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
func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	err := ValidateToken(w, r)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		users := model.GetAllUser()
		json.NewEncoder(w).Encode(users)
	}
}
