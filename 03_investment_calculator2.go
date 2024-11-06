package main

import (
	"fmt"
	"math"
)

// Investment calculator with explicit type assignment
// withExplicitAssignment

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	const inflationRate = 2.5
	// using the const variable
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
