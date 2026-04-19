package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct { // structs are passed by values
	// if the prop is lower case it is private
	// if the prop is upper case it is public
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// struct embedding

type Admin struct {
	email    string //private
	password string //private
	User            // public and that way the methods of the embedded struct are directly accessible
}

// method from User struct
// user is a receiver, a receiver is a copy
func (user User) PrintUser() { //acessing props from structs address do not demand dereferencing it
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

// constructor
func New(firstName string, lastName string, birthdate string) (*User, error) { //for avoiding unnecessary copies
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}
	return &User{firstName: firstName, lastName: lastName, birthdate: birthdate, createdAt: time.Now()}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	user, err := New("ADMIN", "ADMIN", "---")
	if err != nil {
		return nil, err
	}

	return &Admin{
		email:    email,
		password: password,
		User:     *user,
	}, nil
}

// receiver here so should accept the pointer as to alter the content by the reference
func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}
