package main

// Pointers to avoid unnecessary value copies

import (
	"fmt"
)

func avoidCopies() {
	age := 40
	// agePointer := &age   //output= 0xc00011a040

	fmt.Println("Age: ", age)
	fmt.Println("Adult Years: ", getAdultYears(&age))
	fmt.Println("Age: ", age)
}

func getAdultYears(age *int) int {
	return *age - 18 // it will not mutate original age value
}
