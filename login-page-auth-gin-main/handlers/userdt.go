package handlers

import (
	"fmt"

	"main.go/models"
)

var Users = make(map[string]models.User)

func GetUser(Key string) (u models.User, err any) {
	u, exists := Users[Key]
	if !exists {
		err = "User don't exists"
		return u, err
	}
	err = nil
	return u, err
}

func MakeUser(name, email, pwd string) (err any) {
	_, exists := Users[email]
	if exists {
		fmt.Println("User already exists:", Users[email])
		return " already exists"
	}
	Users[email] = models.User{
		Name:     name,
		Email:    email,
		Password: pwd,
	}
	fmt.Println("user created:", Users[email])
	return nil

}
