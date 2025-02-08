package conversation

import (
	"errors"
	"strconv"
)

func Conversation(lines []string) ([]float64, error) {
	stringToFloat := make([]float64, len(lines))
	for index, line := range lines {
		price, error := strconv.ParseFloat(line, 64)
		if error != nil {
			return nil, errors.New("converting failed")
		}
		stringToFloat[index] = price

	}
	return stringToFloat, nil

}
