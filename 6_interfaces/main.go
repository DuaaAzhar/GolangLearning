// Interfaces
// embedded interfaces
// type assertion by interfaces using interface{}
// type switches
// Extracting type information
// Dynamic types with interfaces
// Generics with intefaces

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

// embed the interface in other interface
type outputtable interface {
	saver
	displayer
}

func main() {
	// title, content, err := getNoteData()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// myNote := note.New(title, content)

	// Taking user
	// text, err := getLongUserInput("Todo Text: ")

	// if err != nil {
	// 	fmt.Println("Invalid Todo Text!")
	// 	return
	// }
	// todoText, err := todo.New(text)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// err = outputData(myNote)
	// if err != nil {
	// 	return
	// }
	// outputData(todoText)

	// printAnyType(1)
	// printAnyType(0.1)
	// printAnyType(myNote)

	// printSomeType(1)
	// printSomeType(0.9)
	// printSomeType(myNote)

	result := addGeneric2(0.1, 0.1)
	fmt.Println(result)

}

func getNoteData() (string, string, error) {
	title, err := getLongUserInput("Note Title: ")

	if err != nil {
		return "", "", err
	}

	content, err := getLongUserInput("Note Content: ")
	if err != nil {
		return "", "", err
	}
	return title, content, nil
}

func getLongUserInput(prompt string) (string, error) {
	fmt.Printf("%v", prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text, nil

}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	fmt.Println("Saved Successfully!!")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func printAnyType(data interface{}) {
	fmt.Println(data)
}

// will handle only selected types and not through error if passed other types
func printSomeType(data interface{}) {
	// switch data.(type) {
	// case int:
	// 	fmt.Println(data)
	// case float32:
	// 	fmt.Println(data)
	// case string:
	// 	fmt.Println(data)
	// }

	//  Or

	intValue, ok := data.(int)
	if ok {
		fmt.Println(intValue)
		return
	}

	floatValue, ok := data.(int)
	if ok {
		fmt.Println(floatValue)
		return
	}

	stringValue, ok := data.(string)
	if ok {
		fmt.Println(stringValue)
		return
	}

}

// Dynamic types with interfaces
func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	aString, aIsString := a.(float64)
	bString, bIsString := b.(float64)

	if aIsString && bIsString {
		return aString + bString
	}
	return nil
}

// to avoid the repition of code as in add(), we will introduce generics
// in this case go will predict the return types when that function is called, i.e. float64 in this case when called

// func addGeneric1[T any](a, b T) T {
// 	return a + b
// }

// but to avoid this error of + operator not defined on variable of type T, we can also give optional types to T

func addGeneric2[T int | float64 | string](a, b T) T {
	return a + b
}
