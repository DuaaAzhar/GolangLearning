// make function for arrays and maps
// type aliases
// for loops with maps and arrays
package main

import (
	"fmt"
)

// to avoid to write long types repeatedly
type floatMap map[string]float64

func Output(m floatMap) {
	fmt.Println(m)
}

func main() {
	// 2= length of array , 5= capacity of array
	userNames := make([]string, 2, 5)
	fmt.Println(userNames)

	// will not utilize first two index
	userNames = append(userNames, "Duaa", "Azhar")
	fmt.Println(userNames)

	userNames[0] = "My"
	userNames[1] = "Name"
	fmt.Println(userNames)
	for index, value := range userNames {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}

	// using make for maps
	// in this case only 2 args can be passed

	courseRatings := make(floatMap, 3)
	courseRatings["Go"] = 4.1
	courseRatings["React"] = 4.5
	courseRatings["Javascript"] = 4.7

	// it will re allocate the memory in this case and increase the size
	courseRatings["Nodejs"] = 4.7

	Output(courseRatings)
	for index, value := range courseRatings {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}

}
