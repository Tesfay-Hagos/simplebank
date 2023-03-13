package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tesfayprep/user_authentication/model"

	"golang.org/x/crypto/bcrypt"
)

var Totp = *OtpGenerator()

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
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid body")
		return
	}
	dbuser := model.GetUser(user.Email)
	equal := bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(user.Password))
	if equal != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "can not authenticate this user")
		return
	} else {
		token, err := GenerateJWT(user.Username)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error in generating token")

		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, token)
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
func ResetPassRequestNew(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Start Reseting password")
	user := model.ResetPassword{}
	err := json.NewDecoder(r.Body).Decode(&user)
	model.CheckErr(err)
	if user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please insert both Email to reset password")
	}
	userdb := model.GetUser(user.Email)
	if userdb.Email == user.Email {
		secritcode := Totp.Get()
		//model.SendEmail("tesfay.hagos1421@gmail.com", "TsadkanBerut2121@Adigrat", user.Email, "RESET PASSWORD TOKEN", fmt.Sprintf("This is your Reset Password token:%s", secritcode))
		json.NewEncoder(w).Encode(fmt.Sprint(secritcode))

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func OtpResetpass(w http.ResponseWriter, r *http.Request) {
	data := model.ResetPasswordform{}
	err := json.NewDecoder(r.Body).Decode(&data)
	model.CheckErr(err)
	if data.Token == "" || data.NewPassword == "" || data.ConfirmNewPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please insert token,newpassword and confirmnewpasssword")
	}
	pass := Totp.Verify(data.Token)
	fmt.Printf("OtpCode Received:%s", data.Token)
	if pass {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "This token is already out of date")
	} else {
		model.ChngePassword(data.Email, data.NewPassword)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Password Updated Successfully")
	}

}
