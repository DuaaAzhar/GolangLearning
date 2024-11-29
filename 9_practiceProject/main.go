package main

import (
	cmdmanager "github.com/practiceProject/cmdManager"
	"github.com/practiceProject/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("taxPrice_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		job.Process()

	}
}
