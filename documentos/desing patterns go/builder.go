package main

import (
	"fmt"
)

type User struct {
	Name     string
	Email    string
	Age      int
	Phone    string
	Address  string
	Verified bool
}

func NewUser(name string, email string, age int, phone string, address string, verified bool) User {
	return User{
		Name:     name,
		Email:    email,
		Age:      age,
		Phone:    phone,
		Address:  address,
		Verified: verified,
	}
}

func main() {
	user := NewUser("Alice", "alice@example.com", 0, "", "", false)
	fmt.Printf("%+v\n", user)
}
