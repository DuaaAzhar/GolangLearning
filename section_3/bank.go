// Exercise to deal with error handling and file handling
//  1) Validate user input
//   => Show error msg and exit if invalid input is provided
//   => No negative numbers
//   => No zero valueRequired
//   => Store calculated results in file

package main

import (
	"fmt"

	"os"
)

const taxRate float64 = 10

func mainErrorHandling() {
	var revenue float64
	var expenses float64 = 1000

	revenue = takeInput("Enter Total Revenue: ")
	expenses = takeInput("Enter Total Expenses: ")

	EBT, profit, ratio := profitCalculator(revenue, expenses)
	result := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", EBT, profit, ratio)
	fmt.Println(result)

	writeToFile(result)

}

func takeInput(valueRequired string) float64 {
	var input float64
	fmt.Print(valueRequired)
	fmt.Scanln(&input)
	if input <= 0 {
		panic("0 or -ve Values Not Allowed!! ")
	}
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

func writeToFile(result string) {
	os.WriteFile("data.txt", []byte(result), 0644)

	// ..........IF WE WANT APPEND OPERATION...........
	// file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	// if err != nil {
	// 	fmt.Println("Failed to open file")
	// 	return
	// }
	// defer file.Close()

	// _, err2 := file.WriteString(fmt.Sprint(data))
	// if err2 != nil {
	// 	fmt.Println("Failed to write on file")
	// } else {
	// 	fmt.Println("Successfully updated the file")
	// }
}
