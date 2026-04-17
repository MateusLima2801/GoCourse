package main

import (
	"fmt"
	"math"
)

func calculateInvestmentAmount() {
	const inflationRate float64 = 2.5
	var investmentAmount, years float64 = 1000.0, 10.0
	expectedReturnRate := 5.5 // inferred type

	fmt.Scan(&investmentAmount) // using pointer for storing value scanned from stdin

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value in", years, "years:", futureValue)
	fmt.Println("Future Real Value in", years, "years:", futureRealValue)
}
