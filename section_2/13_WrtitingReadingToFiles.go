// Reading data from files
// writing data to files
// error handling

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeFile(balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceString), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, errors.New("Failed to find balance file")
	}

	stringData := string(data)
	floatData, err := strconv.ParseFloat(stringData, 64)
	if err != nil {
		return 0, errors.New("Failed to parse balance")
	}
	return floatData, nil
}

func mainFileHandling() {
	var choice int
	AccountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------------")
		// return                       // if you want the program to stop when u get this error
		panic("Cannot continue ...") // instead of return if you want the program to stop with some message and info
	}

	fmt.Println("..........Welcome to Bank!!..............")

	for {
		fmt.Println("Which Service you need?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Amount")
		fmt.Println("3. Withdraw Amount")
		fmt.Println("4. Exit")

		fmt.Println("Enter your Choice")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your Account Balance is %.2f\n\n", AccountBalance)
		case 2:
			fmt.Println("Kindly Enter the Amount to deposit")
			var amountToDeposit float64
			fmt.Scan(&amountToDeposit)
			if amountToDeposit <= 0 {
				fmt.Println("Amount to be deposit should be greater than 0")
				// return    //if we apply return it will exit the whole program
				continue // continue will only exit the current loop and start the next iteration
			}
			AccountBalance += amountToDeposit
			writeFile(AccountBalance)
			fmt.Printf("Your final Balance is %.2f\n\n", AccountBalance)

		case 3:
			fmt.Println("Kindly Enter the Amount to withdraw")
			var amountToWithdraw float64
			fmt.Scan(&amountToWithdraw)

			if amountToWithdraw <= 0 || AccountBalance-amountToWithdraw < 0 {
				fmt.Println("Invalid amount to withdraw")
				continue
			}

			AccountBalance -= amountToWithdraw
			writeFile(AccountBalance)
			fmt.Printf("Your final Balance is %.2f\n\n", AccountBalance)
			// break // no need to add break like of other languages switch

		default:
			fmt.Println("GoodBye!")
			fmt.Println("Thanks for using our service")
			// break //in switches of Go, this break will breaks out to next iteration but will not do exit
			return
		}

	}

}
