package main

import (
	"fmt"

	"example.com/greetings"
)

func callGreetings() {
	// get a greeting msg and print it
	message := greetings.Hello("Duaa")
	fmt.Println(message)
}
