package main

import (
	"fmt"
	"math"
)

// Investment calculator with type conversions

func withTypeConversion() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue) // will start a new line on console

	fmt.Print(futureValue) // will not start a new line on console

}
