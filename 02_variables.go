package main

import (
	"fmt"
)

func variables() {
	// var investmentAmount float64 = 1000
	// if we dont want to explicity assign type, we can ommit var keyword
	// expectedReturnRate := 5.5 // type will be inferred by compiler

	// we can also define multiple variables in single line
	// var a, b, c = 10, "20", true
	// a, b, c := 10, "20", true

	// can also use type cast for multiple declared variables, but can be casted to only one type
	var a, b, c float64 = 10, 20, 30
	fmt.Println(a, b, c)

}
