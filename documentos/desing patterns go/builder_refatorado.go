package main

import "fmt"

type User struct {
	Name     string
	Email    string
	Age      int
	Phone    string
	Address  string
	Verified bool
}

type UserBuilder struct {
	user User
}

func NewUserBuilder(name, enail string) *UserBuilder {
	return &UserBuilder{
		user: User{
			Name:  name,
			Email: enail,
		},
	}
}

func (b *UserBuilder) Age(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) Phone(phone string) *UserBuilder {
	b.user.Phone = phone
	return b
}

func (b *UserBuilder) Address(address string) *UserBuilder {
	b.user.Address = address
	return b
}

func (b *UserBuilder) Verified(verified bool) *UserBuilder {
	b.user.Verified = verified
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}

func main() {
	user := NewUserBuilder("Alice", "alice@example.com").
		Age(30).
		Phone("123-456-7890").
		Verified(true).
		Build()

	fmt.Printf("%+v\n", user)
}
