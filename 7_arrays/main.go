// // arrays
// // slice
// // lists with slices

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	prices := []float64{0.0, 0.1, 0.2, 0.3, 0.4}
// 	var productNames [4]string = [4]string{"pencils", "pens", "sharpeners", "erasers"}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	// slicing the prices array
// 	slicedArray := prices[:3]
// 	fmt.Println(slicedArray)
// 	// output= [0 0.1 0.2]

// 	sliceOfSliced := slicedArray[1:]
// 	fmt.Println(sliceOfSliced)
// 	// output= [0.1 0.2]

// 	fmt.Println(len(slicedArray), cap(sliceOfSliced))
// 	// output = 3 4
// 	fmt.Println(len(sliceOfSliced), cap(sliceOfSliced))
// 	// output = 2 4

// 	// can alter the original array
// 	sliceOfSliced[0] = 5.4
// 	fmt.Println(prices)
// 	// output= [0 5.4 0.2 0.3 0.4]

// 	// slices creating the dynamic arrays as they allow the operation of append and delete

// 	// appending element
// 	numbers := []int{1, 2, 3, 4}
// 	sliceOfNumbers := numbers[:2]
// 	appendedArray := append(numbers, 6)
// 	fmt.Println("appended array: ", appendedArray)
// 	fmt.Println("numbers array: ", numbers) //will not change the original array

// 	fmt.Println("Slice of Numbers: ", sliceOfNumbers)

// 	appendedNumbers := append(sliceOfNumbers, 5)
// 	fmt.Println("appended numbers: ", appendedNumbers)
// 	fmt.Println("numbers array: ", numbers) //appending to the slice of array will alter the original array

// 	// deleting element
// 	numbers = numbers[:3]
// 	fmt.Println("numbers array after deleting last 4th index: ", numbers)

// }
