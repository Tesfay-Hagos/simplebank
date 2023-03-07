package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tesfayprep/user_authentication/model"
	"time"

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
			fmt.Fprint(w, token)
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
	user := model.UserInfo{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		users := model.GetAllUser()
		json.NewEncoder(w).Encode(users)
	}
}
func ResetpasswordRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Start Reseting password")
	user := model.RestPassword{}
	err := json.NewDecoder(r.Body).Decode(&user)
	model.CheckErr(err)
	if user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please insert both Email to reset password")
	}
	userdb, isfound := model.GetUser(user.Email)
	if isfound && userdb.Email == user.Email {
		secritcode := GenerateResetToken(64)
		//GenerateJWT(userdb.Username)
		storedtoken := model.Resetpassworddb{UserID: user.Email, Token: secritcode, TokenExpiry: time.Now().Add(time.Minute * 1).Unix()}
		response := model.Inserttoken(storedtoken)
		//model.SendEmail("tesfay.hagos1421@gmail.com", "TsadkanBerut2121@Adigrat", user.Email, "RESET PASSWORD TOKEN", fmt.Sprintf("This is your Reset Password token:%s", secritcode))
		json.NewEncoder(w).Encode(response)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	//todo
}
func ResetPasswordWithToken(w http.ResponseWriter, r *http.Request) {
	data := model.ResetPasswordform{}
	err := json.NewDecoder(r.Body).Decode(&data)
	model.CheckErr(err)
	if data.Token == "" || data.NewPassword == "" || data.ConfirmNewPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please insert token,newpassword and confirmnewpasssword")
	}
	tokens := model.ReadResetTokens(data.Email)
	tokenfound := false
	tokenvalid := false
	if len(tokens) != 0 {
		for _, val := range tokens {
			if errv := ValidateResetToken(val.Token, data.Token); errv == nil {
				tokenvalid = true
				if val.UserID == data.Email && data.NewPassword == data.ConfirmNewPassword {
					tokenfound = true
					model.ChngePassword(data.Email, data.NewPassword)
					w.WriteHeader(200)
					fmt.Fprintf(w, "Passwor Updated Successfully")
				}
			}
		}
		if !tokenfound || !tokenvalid {
			w.WriteHeader(http.StatusExpectationFailed)
			fmt.Fprintf(w, "This token is valid")
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "This token is already out of date no tokens found")
	}

}

/*
func hashpassword(pw []byte) []byte {
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
*/
