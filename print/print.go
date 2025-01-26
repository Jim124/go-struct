package print

import "fmt"

func Print(value any) {
	switch value.(type) {
	case int:
		fmt.Println("i'm a int value")
	case float64:
		fmt.Println("I'm a float value")
	case string:
		fmt.Println("I'm a string value")
	default:
		fmt.Println("choose another type")
	}

}
func PrintWithType(value any) {
	typedVale, ok := value.(int)
	if ok {
		fmt.Printf("I am a int value %v\n", typedVale)
	}
	floatValue, ok := value.(float64)
	if ok {
		fmt.Printf("I am a float value %v\n", floatValue)
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Printf("I am a string value %v\n", stringVal)
	}
	boolVal, ok := value.(bool)
	if ok {
		fmt.Printf("I am a bool value %v\n", boolVal)
	}
}
