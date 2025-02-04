package price

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	file, error := os.Open("prices.txt")
	if error != nil {
		fmt.Println(error)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	error = scanner.Err()
	if error != nil {
		fmt.Println(error)
		return
	}
	prices := make([]float64, len(lines))
	for index, line := range lines {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("converting price to float failed")
			return
		}
		prices[index] = price
	}
	job.Prices = prices
}

func (job TaxIncludePrice) Process() {
	result := make(map[string]string)
	for _, price := range job.Prices {
		priceStr := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = priceStr
	}
	fmt.Println(result)
}
