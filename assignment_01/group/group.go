package group

import "time"

// For users represented by []User struct, group the users by their week number
type User struct {
	ID      string
	Created time.Time
}
type CreatedWeek struct {
	WeekNumber int
	Users      []User
}

func GroupUsersByWeek(users []User) []CreatedWeek {
	// TODO: add your code here
	return []CreatedWeek{}
}
