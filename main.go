package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-struct/array"
	"github.com/go-struct/embed"
	"github.com/go-struct/fileManage"
	"github.com/go-struct/funcVal"
	"github.com/go-struct/map_tutorial"
	"github.com/go-struct/price"
	"github.com/go-struct/print"
	"github.com/go-struct/project"
	"github.com/go-struct/recursion"
	"github.com/go-struct/saver"
	"github.com/go-struct/sum"
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
	//make.TestMakeSLice()
	//make.TestMakeMap()
	funcVal.TestFuncVal()
	numbers := recursion.Factorial(5)
	fmt.Println(numbers)
	sum.TestSum()
	project.Test()

	taxRates := []float64{0.12, 0.23, 0.56}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {
		// cmd := cmdManage.New()
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := fileManage.New("prices.txt", fmt.Sprintf("result-%.1f.json", taxRate))
		taxIncludePrice := price.NewTaxIncludePrice(fm, taxRate)
		go taxIncludePrice.Process(doneChans[index], errorChans[index])
		// if error != nil {
		// 	fmt.Println(error)
		// }
	}
	// for index := range doneChans {
	// 	<-doneChans[index]
	// }
	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			fmt.Println(err)
		case done := <-doneChans[index]:

			fmt.Println(done)
		}

	}
	// concurrency.Test()
	// multipleChan.TestMultipleChan()
	//oneChan.TestOneChan()

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
