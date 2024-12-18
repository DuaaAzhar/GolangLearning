package main

import "fmt"

// Time to practice what you learned!

//  1. Create a new array (!) that contains three hobbies you have
//     Output (print) that array in the command line.
//  2. Also output more data about that array:
//     - The first element (standalone)
//     - The second and third element combined as a new list
//  3. Create a slice based on the first element that contains
//     the first and second elements.
//     Create that slice in two different ways (i.e. create two slices in the end)
//  4. Re-slice the slice from (3) and change it to contain the second
//     and last element of the original array.
//  5. Create a "dynamic array" that contains your course goals (at least 2 goals)
//  6. Set the second goal to a different one AND then add a third goal to that existing dynamic array
//  7. Bonus: Create a "Product" struct with title, id, price and create a
//     dynamic list of products (at least 2 products).
//     Then add a third product to the existing list of products.
type product struct {
	title string
	id    int
	price float64
}

func main() {
	// -----------------------------------1-------------------------------
	var hobbies [3]string
	hobbies = [3]string{"Music", "Movies", "Outing"}

	fmt.Println(hobbies)

	// -----------------------------------2-------------------------------
	fmt.Println(hobbies[0])

	fmt.Println(hobbies[1:3])
	// or
	fmt.Println(hobbies[1:])

	// -----------------------------------3-------------------------------
	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)

	// -----------------------------------4-------------------------------
	sliceOfSlice1 := slice1[1:3]
	fmt.Println("slice of slice1: ", sliceOfSlice1)

	// -----------------------------------5-------------------------------
	var goals []string = []string{"Learn Basics", "Practice Go"}
	fmt.Println("Goals: ", goals)

	// -----------------------------------6-------------------------------
	goals = append(goals, "In depth")
	goals[1] = "Practice Golang Code"
	fmt.Println("Goals after append: ", goals)

	// -----------------------------------7-------------------------------

	var products []product = []product{{"Shirts", 1, 100.1}, {"Trousers", 2, 200.2}}
	products = append(products, product{"Jackets", 3, 300.3})
	fmt.Println(products)

}
