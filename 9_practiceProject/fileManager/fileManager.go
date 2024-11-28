package filemanager

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	// Open file
	file, err := os.Open(path)
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
