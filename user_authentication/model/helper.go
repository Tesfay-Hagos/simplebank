package model

import (
	"fmt"
	"net/smtp"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
func PrintMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func SendEmail(from, password, toEmailAddress, subject, body string) {
	to := []string{toEmailAddress}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
}
