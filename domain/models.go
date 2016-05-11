package domain

import "fmt"

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

var UserStore []User

func (u *User) ToString(user interface{}) string {
	return fmt.Sprint(user)
}
