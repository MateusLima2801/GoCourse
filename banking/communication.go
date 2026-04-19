package banking

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func WelcomeMessage() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 - ", randomdata.PhoneNumber())
}

func Menu() int {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	return choice
}
