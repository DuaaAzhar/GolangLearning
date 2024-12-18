package main

import (
	"fmt"

	filemanager "github.com/goroutinesAndChannels/fileManager"
	"github.com/goroutinesAndChannels/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	dones := make([]chan bool, len(taxRates))

	errChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		dones[index] = make(chan bool)
		errChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("taxPrice_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go job.Process(dones[index], errChans[index])

		// if err != nil {
		// 	fmt.Println("Failed to Process job")
		// }

	}

	for index := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-dones[index]:
			fmt.Println("DONE!")
		}
	}
}
