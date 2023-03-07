package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

/*
create TABLE password_reset_tokens (

	user_id VARCHAR(36) NOT NULL,
	token VARCHAR(128) NOT NULL UNIQUE,
	token_expiry BIGINT NOT NULL,
	PRIMARY KEY (user_id, token) );
*/
type UserInfo struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
type Resetpassworddb struct {
	UserID      string
	Token       string
	TokenExpiry int64
}
type ResetPasswordform struct {
	Email              string
	Token              string
	NewPassword        string
	ConfirmNewPassword string
}
type TokenResponsejson struct {
	Type   string
	Mesage string
	Token  string
}
type RestPassword struct {
	Email string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "tesfay2f"
	password = "tsionawi@2121"
	dbname   = "tesfay2fdb"
)

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)

	CheckErr(err)

	return db
}

type JsonResponse struct {
	Type    string     `json:"type"`
	Data    []UserInfo `json:"data"`
	Message string     `json:"message"`
}
