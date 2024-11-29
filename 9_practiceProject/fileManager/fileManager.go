package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	// Open file
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("Failed to open file")
		fmt.Println(err)
		file.Close()
		return nil, err
	}
	// Read Files
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Failed to read file")
		fmt.Println(err)
		file.Close()
		return nil, err
	}
	file.Close()
	return lines, nil

}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON")

	}
	file.Close()
	return nil
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}
