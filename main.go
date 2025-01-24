package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-struct/node"
)

func main() {
	title, content := getNodeData()
	userNode, error := node.New(title, content)
	if error != nil {
		fmt.Println(error)
		return
	}
	userNode.Display()
	userNode.SaveToFile()

}

func getNodeData() (string, string) {
	title := getUserInput("enter node title: ")
	content := getUserInput("enter node content: ")
	return title, content
}

func getUserInput(hint string) string {
	fmt.Printf("%v", hint)
	reader := bufio.NewReader(os.Stdin)
	text, error := reader.ReadString('\n')
	if error != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
