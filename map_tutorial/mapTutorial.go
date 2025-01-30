package map_tutorial

import "fmt"

func TestMap() {
	persons := map[string]int{"jim": 12, "daniel": 23}
	value := persons["jim"]
	fmt.Println(value)
	delete(persons, "jim")
	fmt.Println(persons)
	clear(persons)
	fmt.Println(persons)
}
