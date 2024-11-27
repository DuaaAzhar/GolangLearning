package main

import (
	"fmt"
)

func main() {
	URLS := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	// it prints the values in alphabetical order
	fmt.Println(URLS)
	// Accessing maps
	fmt.Println(URLS["Google"])

	// adding items
	URLS["Github"] = "https://github.com"
	fmt.Println("After Adding items===> ", URLS)

	// deleting items
	delete(URLS, "Amazon Web Services")
	fmt.Println("After Deleting items===> ", URLS)

}
