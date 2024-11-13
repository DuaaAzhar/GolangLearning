package main

import (
	"fmt"
)

func formatStrings() {
	investmentAmount := 100.4579739
	taxRate := 100.0
	// f represents float number.    .0f will represent the no. of decimals we want after point
	fmt.Printf("Investment amount %s we have is of Sales department with tax rate %s\n", investmentAmount, taxRate)
	fmt.Printf("Investment amount %d we have is of Sales department with tax rate %d\n", investmentAmount, taxRate)
	fmt.Printf("Investment amount %.2f we have is of Sales department with tax rate %.0f\n", investmentAmount, taxRate)

}
