package fileManage

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManage struct {
	inputPath  string
	outputPath string
}

func (fm *FileManage) ReadLine() ([]string, error) {
	file, error := os.Open(fm.inputPath)
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

func (fm *FileManage) WriteResult(data any) error {
	file, error := os.Create(fm.outputPath)
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

func New(inputPath, outputPath string) *FileManage {
	return &FileManage{
		inputPath:  inputPath,
		outputPath: outputPath,
	}
}
