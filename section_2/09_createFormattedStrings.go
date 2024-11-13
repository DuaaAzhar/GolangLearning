package main

import (
	"fmt"
)

func createFormattedStrings() {
	investmentAmount := 100.4579739
	taxRate := 100.0
	// instead of formatting and printing numbers to console, we need to store them in variables
	// for this we can use Sprint
	formattedValue := fmt.Sprintf("Investment amount %.2f we have is of Sales department with tax rate %.0f\n", investmentAmount, taxRate)
	fmt.Print(formattedValue)

	// instead of using \n we can add line breaks directly by using back ticks instead of double quotes
	fmt.Printf(`Investment amount %.2f,
	 we have is of Sales department with tax rate %.0f`, investmentAmount, taxRate)
}
