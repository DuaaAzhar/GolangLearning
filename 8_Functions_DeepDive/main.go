// passing functions as params in functions
// anonymous functions

package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	dNumbers := transformNumbers(&numbers, double)
	fmt.Println(dNumbers)
	tNumbers := transformNumbers(&numbers, triple)
	fmt.Println(tNumbers)

	moreNumbers := []int{2, 1, 3, 4}
	func1 := getTransformerFunction(&numbers)
	func2 := getTransformerFunction(&moreNumbers)

	fmt.Println(transformNumbers(&numbers, func1))
	fmt.Println(transformNumbers(&numbers, func2))

	// anonymous functions
	transformedNumbers := transformNumbers(&numbers, func(value int) int {
		return value * 4
	})
	fmt.Println("by anonymous funcs: ", transformedNumbers)

	// closures
	fmt.Println("doubled with closures ", createTransformer(2))
	fmt.Println("doubled with closures ", createTransformer(3))
}

// passing functions as param in functions
func transformNumbers(numbers *[]int, transform transformFn) []int {
	var dNumbers []int
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(value int) int {
	return value * 2
}

func triple(value int) int {
	return value * 3
}

// returning functions
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

// closures

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
