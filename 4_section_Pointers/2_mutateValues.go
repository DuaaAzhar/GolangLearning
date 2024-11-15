package main

// Pointers to avoid unnecessary value copies

import (
	"fmt"
)

func main() {
	age := 40
	// agePointer := &age   //output= 0xc00011a040

	fmt.Println("Before edit Age: ", age)
	editAgeToAdultYears(&age)
	fmt.Println("After edit Age: ", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
