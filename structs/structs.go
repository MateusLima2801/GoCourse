package structs

import (
	"fmt"

	"go.com/structs/user"
)

func StartStructs() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthdate)
	adminUser, err := user.NewAdmin("test@example.com", "***")
	// nullUser := User{}
	if err != nil {
		fmt.Println(err)
		return
	}
	adminUser.PrintUser()
	adminUser.ClearUserName()
	adminUser.PrintUser()
	appUser.PrintUser()
	appUser.ClearUserName()
	appUser.PrintUser()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
