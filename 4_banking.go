package main

import (
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func banking() {
	WelcomeMessage()
	choice := 0
	var balance float64 = ReadFile()
	for choice != 4 {
		choice = Menu()
		HandleMenuChoice(choice, &balance)
	}
}

func WriteToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func ReadFile() float64 {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		os.Create(balanceFile)
		return 0
	}
	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		panic("Invalid entry")
	}
	return balance
}

func HandleMenuChoice(choice int, balance *float64) {
	switch choice {
	case 4:
		fmt.Println("Goodbye, thank you for using Go Bank")
	case 1:
		CheckBalance(*balance)
	case 2:
		TransactMoney(balance, true)
	case 3:
		TransactMoney(balance, false)
	default:
		fmt.Println("Invalid choice, please try again.")
	}
}

func TransactMoney(balance *float64, isDeposit bool) {
	var transact float64
	fmt.Print(getMessage(isDeposit))
	fmt.Scan(&transact)
	if !IsValid(*balance, isDeposit, transact) {
		fmt.Println("Invalid transaction.")
		return
	}
	if isDeposit {
		*balance += transact
	} else {
		*balance -= transact
	}
	fmt.Printf("Balance updated! New amount: $%.2f\n", *balance)
	WriteToFile(*balance)
}

func IsValid(balance float64, isDeposit bool, transact float64) bool {
	if isDeposit {
		return transact > 0
	}
	return transact > 0 && balance >= transact
}

func getMessage(isDeposit bool) string {
	if isDeposit {
		return "Your deposit: "
	}
	return "Your withdrawal: "
}

func CheckBalance(balance float64) {
	fmt.Printf("Your current balance is: $%.2f\n", balance)
}

func WelcomeMessage() {
	fmt.Println("Welcome to Go Bank!")
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
