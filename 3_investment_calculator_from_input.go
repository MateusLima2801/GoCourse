package main

import (
	"fmt"
	"math"
)

func calculateInvestmentAmountFromInput() {
	const inflationRate float64 = 2.5
	var investmentAmount, years, expectedReturnRate float64

	getFloatFromInput("Investment amount", &investmentAmount)
	getFloatFromInput("Expected return rate (in %)", &expectedReturnRate)
	getFloatFromInput("Number of years", &years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate)
	//fmt.Println("Future Value in", years, "years:", futureValue)
	//fmt.Println("Future Real Value in", years, "years:", futureRealValue)
	fmt.Printf("Future Value in %v years: %.2f\n", years, futureValue)
	//fmt.Printf("Future Real Value in %v years: %.2f\n", years, futureRealValue)
	result := fmt.Sprintf("Future Real Value in %v years: %.2f\n", years, futureRealValue)
	fmt.Print(result)
}

func getFloatFromInput(name string, variable *float64) {
	fmt.Print(name, ": ")
	fmt.Scan(variable)
}

func calculateRateOverTime(rate float64, years float64) float64 {
	return math.Pow(1+rate/100, years)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64, inflationRate float64) (float64, float64) {
	var futureValue = investmentAmount * calculateRateOverTime(expectedReturnRate, years)
	var futureRealValue = futureValue / calculateRateOverTime(inflationRate, years)
	return futureValue, futureRealValue
}
