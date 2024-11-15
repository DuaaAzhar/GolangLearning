package main

import (
	"errors"
	"fmt"

	"github.com/practice/note"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}

	myNote := note.New(title, content)

}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title: ")

	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note Content: ")
	if err != nil {
		return "", "", err
	}
	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	if value != "" {
		return "", errors.New("invalid input")
	}
	return value, nil
}
