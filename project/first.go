package project

import "fmt"

func Test() {
	prices := []float64{10, 20, 30}
	taxRate := []float64{0.0, 0.1, 0.2, 0.3}
	taxPrice := make(map[float64][]float64)
	for _, rate := range taxRate {
		taxIncludePrices := make([]float64, len(prices))
		for index, price := range prices {
			taxIncludePrices[index] = price * (1 + rate)
		}
		taxPrice[rate] = taxIncludePrices
	}
	fmt.Println(taxPrice)

}
