package model

import (
	"fmt"
	"log"

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
	response := JsonResponse{Type: "success", Data: users}
	return response
}

func hashpassword(pw []byte) []byte {
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
