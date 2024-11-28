package prices

import (
	"fmt"

	"github.com/practiceProject/conversion"
	filemanager "github.com/practiceProject/fileManager"
)

type TaxIncludedPriceJob struct {
	inputPrices []float64
	taxRate     float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		inputPrices: []float64{},
		taxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, inputPrice := range job.inputPrices {
		taxIncludedPrice := inputPrice * (1 + job.taxRate)
		formattedPrice := fmt.Sprintf("%.2f", inputPrice)
		formattedTaxRate := fmt.Sprintf("%.2f", taxIncludedPrice)
		result[formattedPrice] = formattedTaxRate

	}
	fmt.Println(result)

}

func (job *TaxIncludedPriceJob) LoadData() {

	// Reading lines
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Converting to Float
	job.inputPrices, err = conversion.ConvertStringToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

}
