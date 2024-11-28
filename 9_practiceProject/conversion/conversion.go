package conversion

import (
	"errors"
	"strconv"
)

func ConvertStringToFloat(texts []string) ([]float64, error) {
	var floats []float64
	for _, text := range texts {
		float, err := strconv.ParseFloat(text, 64)
		floats = append(floats, float)
		if err != nil {
			return nil, errors.New("Failed to convert string to float")
		}

	}
	return floats, nil
}
