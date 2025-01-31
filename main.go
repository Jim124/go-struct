package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-struct/array"
	"github.com/go-struct/embed"
	"github.com/go-struct/funcVal"
	"github.com/go-struct/make"
	"github.com/go-struct/map_tutorial"
	"github.com/go-struct/print"
	"github.com/go-struct/saver"
)

func main() {
	// title, content := getNodeData()
	// userNode, error := node.New(title, content)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	// //userNode.Display()
	// //userNode.SaveToFile()
	// play(userNode)
	// text := getUserInput("Please enter what you want to do: ")
	// userTodo, error := todo.New(text)
	// if error != nil {
	// 	fmt.Println(error)
	// }
	// save(userTodo)

	// print something
	print.Print(1)
	print.Print(1.0)
	print.Print("hello")
	// print.Print(userNode)
	print.PrintWithType(1)
	print.PrintWithType(true)
	array.TestArray()
	map_tutorial.TestMap()
	make.TestMakeSLice()
	make.TestMakeMap()
	funcVal.TestFuncVal()

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

func save(data saver.Saver) {
	data.SaveToFile()
}

func play(data embed.Embed) {
	data.Display()
	data.SaveToFile()
}
