package price

import (
	"fmt"

	"github.com/go-struct/conversation"
	fileManage "github.com/go-struct/filemanage"
)

type TaxIncludePrice struct {
	TaxRate float64
	Prices  []float64
}

func NewTaxIncludePrice(taxRate float64) *TaxIncludePrice {
	return &TaxIncludePrice{
		TaxRate: taxRate,
		Prices:  []float64{},
	}
}

func (job *TaxIncludePrice) LoadData() {
	lines, error := fileManage.LoadData("prices.txt")
	if error != nil {
		return
	}
	prices := conversation.Conversation(lines)
	job.Prices = prices
}

func (job *TaxIncludePrice) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.Prices {
		priceStr := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = priceStr
	}
	// fmt.Println(result)
	fileManage.WriteToFile(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), result)
}
