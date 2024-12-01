package main

import (
	"fmt"

	filemanager "github.com/practiceProject/fileManager"
	"github.com/practiceProject/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("pricessss.txt", fmt.Sprintf("taxPrice_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := job.Process()
		if err != nil {
			fmt.Println("Failed to Process job")
		}

	}
}
