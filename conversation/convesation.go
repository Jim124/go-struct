package conversation

import (
	"strconv"
)

func Conversation(lines []string) []float64 {
	stringToFloat := make([]float64, len(lines))
	for index, line := range lines {
		price, error := strconv.ParseFloat(line, 64)
		if error != nil {
			return nil
		}
		stringToFloat[index] = price

	}
	return stringToFloat

}
