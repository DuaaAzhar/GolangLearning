// take revenue, expenses, taxRate from user

// calculate earnings before tax (EBT)
//  calculate earnings after tax (profit)
// calculate ratio (EBT/profit)
// output 3 values EBT, profit and ratio

package main

import (
	"fmt"
)

func ProfitCalculator() {
	var revenue float64 = 2000
	var expenses float64 = 1000
	var taxRate float64 = 10

	fmt.Print("Enter Total Revenue: ")
	fmt.Scanln(&revenue)

	fmt.Print("Enter Total Expenses: ")
	fmt.Scanln(&expenses)

	var taxValue = taxRate / 100.0

	EBT := revenue - expenses
	profit := EBT * (1 - taxValue)
	ratio := EBT / profit

	fmt.Println("EBT:", EBT)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)

}
