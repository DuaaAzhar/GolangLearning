package prices

import (
	"fmt"
)

type TaxIncludedPriceJob struct {
	inputPrices       []float64
	taxRate           float64
	taxIncludedPrices map[string][]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		inputPrices:       []float64{10, 20, 30},
		taxRate:           taxRate,
		taxIncludedPrices: make(map[string][]float64),
	}
}

func (job TaxIncludedPriceJob) Process() {

	for _, inputPrice := range job.inputPrices {
		taxIncludedPrice := inputPrice * (1 + job.taxRate)
		formattedTaxRate := fmt.Sprintf("%.2f", taxIncludedPrice)
		job.taxIncludedPrices[formattedTaxRate] = append(job.taxIncludedPrices[formattedTaxRate], taxIncludedPrice)

	}
	fmt.Println(job)

}
