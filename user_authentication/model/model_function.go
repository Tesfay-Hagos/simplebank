package model

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Register(user UserInfo) JsonResponse {
	var response = JsonResponse{}
	if user.Username == "" || user.Email == "" || user.Password == "" {
		response = JsonResponse{Type: "error", Message: "You are missing non null parameters of the user."}
	} else {
		db := setupDB()
		PrintMessage("Registering user into DB")
		fmt.Println("Inserting new User with Email: " + user.Email + " and name: " + user.Username)
		var lastInsertID int
		err := db.QueryRow("INSERT INTO registeredusers(username,password,email,created_at) VALUES($1,$2,$3,$4) returning id;", user.Username, string(hashpassword([]byte(user.Password))), user.Email, user.CreatedAt).Scan(&lastInsertID)
		// check errors
		CheckErr(err)
		response = JsonResponse{Type: "success", Message: "The  user has been inserted successfully!"}
	}
	return response

}
func GetUser(email string) (UserInfo, bool) {
	//needs to be replaces using Database
	db := setupDB()
	PrintMessage("Getting Users...")
	// Get all users from users table that don't have userID = "1"
	rows, err := db.Query("SELECT * FROM registeredusers")
	// check errors
	CheckErr(err)
	// Foreach user
	for rows.Next() {
		user := UserInfo{}
		var id int
		err = rows.Scan(&id, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if email == user.Email {
			CheckErr(err)
			return user, true
		}
	}
	return UserInfo{}, false
}
func ChngePassword(email, pass string) string {
	db := setupDB()
	PrintMessage("Getting Users...")
	// Get all users from users table that don't have userID = "1"
	sqlStatement := `
	UPDATE registeredusers 
	SET password= $1 WHERE email=$2;`
	_, err := db.Exec(sqlStatement, string(hashpassword([]byte(pass))), email)
	CheckErr(err)
	return "userpassword updated"
}
func GetAllUser() JsonResponse {
	//needs to be replaces using Database
	db := setupDB()
	PrintMessage("Getting Users...")
	users := []UserInfo{}
	response := JsonResponse{}
	// Get all users from users table that don't have userID = "1"
	rows, err := db.Query("SELECT * FROM registeredusers")
	// check errors
	CheckErr(err)
	// Foreach user
	for rows.Next() {
		user := UserInfo{}
		var id int
		err = rows.Scan(&id, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		users = append(users, user)
	}
	CheckErr(err)
	response = JsonResponse{Type: "success", Data: users}
	return response
}
func Inserttoken(t Resetpassworddb) TokenResponsejson {
	response := TokenResponsejson{}
	if t.Token == "" || t.UserID == "" || t.TokenExpiry == 0 {
		response = TokenResponsejson{Type: "error", Mesage: "You are missing non null parameters of the user."}
	} else {
		db := setupDB()
		PrintMessage("Registering Resetpassword")
		//var lastInsertID int
		_, err := db.Exec("INSERT INTO password_reset_tokens(user_id,token,token_expiry) VALUES($1,$2,$3);", t.UserID, hashpassword([]byte(t.Token)), t.TokenExpiry)
		// check errors
		CheckErr(err)
		response = TokenResponsejson{Type: "success", Mesage: fmt.Sprintf("The  token has been send successfully to your email addrest:%s !", t.UserID), Token: t.Token}

	}
	return response
}
func ReadResetTokens(Email string) []Resetpassworddb {
	db := setupDB()
	PrintMessage("Getting Tokens...")
	tokens := []Resetpassworddb{}
	// Get all users from users table that don't have userID = "1"
	rows, err := db.Query("SELECT * FROM password_reset_tokens where user_id=$1", Email)
	// check errors

	CheckErr(err)
	// Foreach user
	for rows.Next() {
		user := Resetpassworddb{}
		//var id int
		err = rows.Scan(&user.UserID, &user.Token, &user.TokenExpiry)
		cond := user.TokenExpiry < time.Now().Unix()
		if cond {
			_, err := db.Exec("DELETE FROM password_reset_tokens where token_expiry=$1", user.TokenExpiry)
			CheckErr(err)
			continue
		}
		tokens = append(tokens, user)
		CheckErr(err)
	}
	CheckErr(err)
	return tokens
}
func hashpassword(pw []byte) []byte {
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
