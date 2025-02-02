package sum

import "fmt"

func sum(number int, numbers ...int) int {
	sum := number
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func TestSum() {
	total := sum(1, 2, 4, 5, 6, 7)
	fmt.Println(total)
	numbers := []int{1, 2, 5, 7, 8}
	total = sum(1, numbers...)
	fmt.Println(total)
}
