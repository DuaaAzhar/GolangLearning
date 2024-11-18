package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/practice/note"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}

	myNote := note.New(title, content)
	myNote.Display()
	errSave := myNote.Save()

	if errSave != nil {
		fmt.Println("Saving to File Failed!!")
	} else {
		fmt.Println("Saving to File Succeeded!!")
	}

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

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	if value == "" {
		return "", errors.New("invalid input")
	}
	return value, nil
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
