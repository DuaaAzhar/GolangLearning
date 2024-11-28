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
	doubled := createTransformer(2)
	tripled := createTransformer(3)
	fmt.Println("doubled with closures ", transformNumbers(&numbers, doubled))
	fmt.Println("doubled with closures ", transformNumbers(&numbers, tripled))

	// recursions
	fmt.Println(factorial(5))

	// Variadic functions
	fmt.Println("sum up of numbers= ", sumUp(1, 2, 3, 4))
	fmt.Println("sum up of numbers= ", sumUp2(1, 2, 3, 4))

	// now instead of writing the array in separated form, if we want to pass the array as separate entities
	fmt.Println(sumUp2(numbers[0], numbers...))
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

// recursions

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// Variadic functions

func sumUp(numbers ...int) int {
	var sum int = 0
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	return sum
}

func sumUp2(num1 int, numbers ...int) int {
	var sum int = 0
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	return sum - num1
}
