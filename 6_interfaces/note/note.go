package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// We will use the struct tags to enable some packages to use them when needed as json package do
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) Note {
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func (note Note) Display() {
	fmt.Printf("Content of Your Note Title: %v has following content:\n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	// Setting the fileName with note's title.
	// Formatting the note title for fileName
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// Storing the data in file in JSON format
	json, err := json.Marshal(note) // we can pass whole struct as param to be converted to json
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
