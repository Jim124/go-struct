package fileManage

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func LoadData(path string) ([]string, error) {
	file, error := os.Open(path)
	if error != nil {
		return nil, errors.New("could not open the file")
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	error = scanner.Err()
	if error != nil {
		return nil, error
	}
	return lines, nil
}

func WriteToFile(path string, data any) error {
	file, error := os.Create(path)
	if error != nil {
		return error
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	error = encoder.Encode(data)
	if error != nil {
		return error
	}
	return nil

}
