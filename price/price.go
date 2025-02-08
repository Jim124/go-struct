package price

import (
	"fmt"

	"github.com/go-struct/conversation"
	"github.com/go-struct/ioManage"
)

type TaxIncludePrice struct {
	TaxRate    float64           `json:"tax_rate"`
	Prices     []float64         `json:"prices"`
	ioManager  ioManage.IOManage `json:"-"`
	TaxInclude map[string]string `json:"taxInclude"`
}

func NewTaxIncludePrice(ioManage ioManage.IOManage, taxRate float64) *TaxIncludePrice {
	return &TaxIncludePrice{
		TaxRate:   taxRate,
		Prices:    nil,
		ioManager: ioManage,
	}
}

func (job *TaxIncludePrice) LoadData() error {
	lines, error := job.ioManager.ReadLine()
	if error != nil {
		return error
	}
	prices, error := conversation.Conversation(lines)
	if error != nil {
		return error
	}
	job.Prices = prices
	return nil
}

func (job *TaxIncludePrice) Process() error {
	error := job.LoadData()
	if error != nil {
		return error
	}
	result := make(map[string]string)
	for _, price := range job.Prices {
		priceStr := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = priceStr
	}

	// fmt.Println(result)
	job.TaxInclude = result
	return job.ioManager.WriteResult(job)
}
