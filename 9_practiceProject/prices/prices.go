package prices

import (
	"fmt"

	"github.com/practiceProject/conversion"
	"github.com/practiceProject/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	InputPrices       []float64           `json:"input_price"`
	TaxRate           float64             `json:"tax_rate"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, inputPrice := range job.InputPrices {
		taxIncludedPrice := inputPrice * (1 + job.TaxRate)
		formattedPrice := fmt.Sprintf("%.2f", inputPrice)
		formattedTaxRate := fmt.Sprintf("%.2f", taxIncludedPrice)
		result[formattedPrice] = formattedTaxRate

	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)

}

func (job *TaxIncludedPriceJob) LoadData() {

	// Reading lines
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Converting to Float
	job.InputPrices, err = conversion.ConvertStringToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

}
