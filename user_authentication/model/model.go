package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type UserInfo struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
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
