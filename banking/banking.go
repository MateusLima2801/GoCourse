package banking

import (
	"fmt"

	"go.com/banking/file"
)

const balanceFile = "banking/balance.txt"

func StartBank() {
	WelcomeMessage()
	choice := 0
	var balance float64 = file.ReadFloatFromFile(balanceFile)
	for choice != 4 {
		choice = Menu()
		HandleMenuChoice(choice, &balance)
	}
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
	file.WriteFloatToFile(*balance, balanceFile)
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
