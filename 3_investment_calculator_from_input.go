package main

import (
	"fmt"
	"math"
)

func calculateInvestmentAmountFromInput() {
	const inflationRate float64 = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount) // using pointer for storing value scanned from stdin
	fmt.Print("Expected return rate (in %): ")
	fmt.Scan(&expectedReturnRate) // using pointer for storing value scanned from stdin
	fmt.Print("Number of years: ")
	fmt.Scan(&years) // using pointer for storing value scanned from stdin

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value in", years, "years:", futureValue)
	fmt.Println("Future Real Value in", years, "years:", futureRealValue)
}
