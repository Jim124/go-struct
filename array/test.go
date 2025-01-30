package array

import "fmt"

type Product struct {
	id    int32
	title string
	price float64
}

func TestArray() {
	hobbies := [3]string{"cooking", "reading", "writing"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	fmt.Println(hobbies[1:3])
	hobbiesSLice := hobbies[:2]
	hobbiesSLice = hobbies[0:2]
	hobbiesSLice = hobbies[1:]
	fmt.Println(hobbiesSLice)
	goals := []string{"go", "react"}
	goals[1] = "javascript"
	goals = append(goals, "typescript")
	fmt.Println(goals)

	var products = []Product{Product{
		id:    1,
		title: "learn array",
		price: 1.0,
	}, Product{
		id:    2,
		title: "learn slice",
		price: 2.0,
	}}
	products = append(products, Product{id: 3, title: "learn typescript", price: 3.0})
	fmt.Println(products)

}
