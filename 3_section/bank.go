// Reading data from files
// writing data to files
// error handling

package main

import (
	"fmt"
	// importing and using self created package
	"example.com/usePackages/fileOps"
	// importing third party package which are not either standard library part nor self created
	"github.com/Pallinder/go-randomdata"
)

func main() {
	var choice int
	fileName := "balance.txt"
	AccountBalance, err := fileOps.GetFloatFromFile(fileName)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------------")
		// return                       // if you want the program to stop when u get this error
		// panic("Cannot continue ...") // instead of return if you want the program to stop with some message and info
	}

	fmt.Println("..........Welcome to Bank!!..............")
	fmt.Println("Contact Us 24/7 at ", randomdata.PhoneNumber())
	for {
		presentOptions()
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
			fileOps.WriteFloatToFile(fileName, AccountBalance)
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
			fileOps.WriteFloatToFile(fileName, AccountBalance)
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
