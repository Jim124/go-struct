package make

import "fmt"

// type ali
type stringMap map[string]int

func TestMakeSLice() {
	names := make([]string, 2, 5)
	fmt.Println(names)
	names[0] = "Jim"
	names = append(names, "Daniel")
	fmt.Println(names)
	for index, value := range names {
		fmt.Println(index, value)
	}
}

func TestMakeMap() {
	person := make(stringMap)
	fmt.Println(person)
	person["jim"] = 35
	person["daniel"] = 28
	fmt.Println(person)
	for k, v := range person {
		fmt.Println(k, v)
	}
}
