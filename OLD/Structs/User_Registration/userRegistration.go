package main

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Email    string
	Password string
}

func (u User) CheckEmail() (bool, error) {

	for i := 0; i < len(u.Email); i++ {
		if u.Email[i] == '@' {
			return true, nil
		}
	}
	return false, errors.New("This is not email, this not has @")
}

func (u User) CheckPassword() (bool, error) {

	if len(u.Password) < 8 {
		return false, errors.New("Password must be at least 8 characters long")
	}
	return true, nil
}

func (u User) HashPassword() string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		errors.New("Failed to hash password")
	}
	return string(hashedPassword)

}

func main() {

	u := User{Username: "John", Email: "john@wp.pl", Password: "123456789"}

	isEmail, err := u.CheckEmail()
	if err != nil {
		println(err.Error())
	}
	println(isEmail)

	isPassword, err := u.CheckPassword()
	if err != nil {
		println(err.Error())
	}
	println(isPassword)

	println(u.HashPassword())

}
