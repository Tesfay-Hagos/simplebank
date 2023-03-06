package model

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func Register(user UserInfo) JsonResponse {
	var response = JsonResponse{}
	if user.Username == "" || user.Email == "" || user.Password == "" {
		response = JsonResponse{Type: "error", Message: "You are missing non null parameters of the user."}
	} else {
		db := setupDB()
		printMessage("Registering user into DB")
		fmt.Println("Inserting new User with Email: " + user.Email + " and name: " + user.Username)
		var lastInsertID int
		err := db.QueryRow("INSERT INTO registeredusers(username,password,email,created_at) VALUES($1,$2,$3,$4) returning id;", user.Username, string(hashpassword([]byte(user.Password))), user.Email, user.CreatedAt).Scan(&lastInsertID)
		// check errors
		checkErr(err)
		response = JsonResponse{Type: "success", Message: "The movie has been inserted successfully!"}
	}
	return response

}
func GetUser(email string) (UserInfo, bool) {
	//needs to be replaces using Database
	db := setupDB()
	printMessage("Getting Users...")
	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM registeredusers")
	// check errors
	checkErr(err)
	// Foreach movie
	for rows.Next() {
		user := UserInfo{}
		var id int
		err = rows.Scan(&id, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if email == user.Email {
			checkErr(err)
			return user, true
		}
	}
	return UserInfo{}, false
}
func GetAllUser() JsonResponse {
	//needs to be replaces using Database
	db := setupDB()
	printMessage("Getting Users...")
	users := []UserInfo{}
	response := JsonResponse{}
	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM registeredusers")
	// check errors
	checkErr(err)
	// Foreach movie
	for rows.Next() {
		user := UserInfo{}
		var id int
		err = rows.Scan(&id, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		users = append(users, user)
	}
	checkErr(err)
	response = JsonResponse{Type: "success", Data: users}
	return response
}
func hashpassword(pw []byte) []byte {
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
