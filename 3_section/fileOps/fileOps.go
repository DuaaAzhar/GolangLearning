package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	valueString := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueString), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("failed to find balance file")
	}

	stringData := string(data)
	floatData, err := strconv.ParseFloat(stringData, 64)
	if err != nil {
		return 0, errors.New("failed to parse balance")
	}
	return floatData, nil
}
