package node

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"encoding/json"
)

type Node struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"createAt"`
}

func New(title, content string) (*Node, error) {
	if title == "" || content == "" {
		return nil, errors.New("invalid input, please check your input value")
	}
	return &Node{
		Title:    title,
		Content:  content,
		CreateAt: time.Now(),
	}, nil
}

func (node Node) Display() {
	fmt.Printf("Node title is: %v\nNode content is: %v\n", node.Title, node.Content)
}

func (node Node) SaveToFile() error {
	fileName := strings.ReplaceAll(node.Title, " ", "-")
	fileName = strings.ToLower(fileName) + ".json"
	json, error := json.Marshal(node)
	if error != nil {
		return error
	}
	return os.WriteFile(fileName, json, 0644)
}
