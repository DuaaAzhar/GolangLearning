package main

import "fmt"

func hello() {

	var student string = "Duaa"
	var student2 = "Azhar" // type is inferred by compiler from its value
	student3 := "Minhas"   // type is inferred by compiler from its value
	fmt.Println(student + " " + student2 + " " + student3)

}
