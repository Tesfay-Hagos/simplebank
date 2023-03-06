package model

import "fmt"

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
