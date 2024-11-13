package main

import (
	"fmt"
)

func mainBankExerciseWithConditions() {
	var choice int
	var AccountBalance float64 = 3000
	// for i := 0; i <= 4; i++ {}  this is controlled loop in which we know how much iterations we need
	// for {}            this is infinite loop
	// for bool(some condition) {}  conditional loop
	fmt.Println("..........Welcome to Bank!!..............")
	for {
		fmt.Println("Which Service you need?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Amount")
		fmt.Println("3. Withdraw Amount")
		fmt.Println("4. Exit")

		fmt.Println("Enter your Choice")
		fmt.Scan(&choice)

		// checkChoice := choice ==1

		if choice == 1 {
			fmt.Printf("Your Account Balance is %.2f\n\n", AccountBalance)

		} else if choice == 2 {
			fmt.Println("Kindly Enter the Amount to deposit")
			var amountToDeposit float64
			fmt.Scan(&amountToDeposit)
			if amountToDeposit <= 0 {
				fmt.Println("Amount to be deposit should be greater than 0")
				// return    //if we apply return it will exit the whole program
				continue // continue will only exit the current loop and start the next iteration
			}
			AccountBalance += amountToDeposit
			fmt.Printf("Your final Balance is %.2f\n\n", AccountBalance)

		} else if choice == 3 {
			fmt.Println("Kindly Enter the Amount to withdraw")
			var amountToWithdraw float64
			fmt.Scan(&amountToWithdraw)

			if amountToWithdraw <= 0 || AccountBalance-amountToWithdraw < 0 {
				fmt.Println("Invalid amount to withdraw")
				continue
			}

			AccountBalance -= amountToWithdraw
			fmt.Printf("Your final Balance is %.2f\n\n", AccountBalance)
		} else if choice == 4 {
			fmt.Println("GoodBye!")
			break // will execute the print statement after ending the loop
			// return //will not execute the print statement after ending the loop
		}

	}
	fmt.Println("Thanks for using our service")

}
