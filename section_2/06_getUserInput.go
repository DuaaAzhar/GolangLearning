package main

import (
	"fmt"
	"math"
)

func getUserInput() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("No. of Years:  ")
	fmt.Scan(&years)

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
