// need to make the general function of taking input from user
// need to make the function for profit calculation

package main

import (
	"fmt"
)

const taxRate float64 = 10

func mainExerciseFunctions() {
	var revenue float64
	var expenses float64 = 1000

	revenue = takeInput("Enter Total Revenue: ")
	expenses = takeInput("Enter Total Expenses: ")

	EBT, profit, ratio := profitCalculator(revenue, expenses)
	fmt.Printf("EBT:  %.1f\n", EBT)
	fmt.Printf("Profit:  %.1f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

}

func takeInput(valueRequired string) float64 {
	var input float64
	fmt.Print(valueRequired)
	fmt.Scanln(&input)
	return input
}

// func profitCalculator(revenue float64, expenses float64) (float64, float64, float64) {
// 	taxValue := taxRate / 100.0
// 	EBT := revenue - expenses
// 	profit := EBT * (1 - taxValue)
// 	ratio := EBT / profit
// 	return EBT, profit, ratio
// }

// other way to return values from a function and pass data types of input params

func profitCalculator(revenue, expenses float64) (EBT float64, profit float64, ratio float64) {
	taxValue := taxRate / 100.0
	EBT = revenue - expenses
	profit = EBT * (1 - taxValue)
	ratio = EBT / profit
	return
}
