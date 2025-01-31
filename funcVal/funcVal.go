package funcVal

import "fmt"

type transformFunc func(int) int

func TestFuncVal() {
	ages := []int{1, 2, 4, 6}
	doubleAges := getTransformVal(ages, getDouble)
	fmt.Println(doubleAges)
	tripleAge := getTransformVal(ages, getTriple)
	fmt.Println(tripleAge)
}

func getTransformVal(ages []int, transform transformFunc) []int {
	values := []int{}
	for _, val := range ages {
		values = append(values, transform(val))
	}
	return values
}

func getDouble(age int) int {
	return age * 2
}
func getTriple(age int) int {
	return age * 3
}
