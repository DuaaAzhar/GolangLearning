package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// We will use the struct tags to enable some packages to use them when needed as json package do
type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println("Your Todo Text: ", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo) // we can pass whole struct as param to be converted to json
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
