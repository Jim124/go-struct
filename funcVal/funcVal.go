package funcVal

import "fmt"

type transformFunc func(int) int

func TestFuncVal() {
	ages := []int{1, 2, 4, 6}
	doubleAges := getTransformVal(ages, getDouble)
	fmt.Println(doubleAges)
	tripleAge := getTransformVal(ages, getTriple)
	fmt.Println(tripleAge)
	// anonymous method
	fourthAge := getTransformVal(ages, func(number int) int { return number * 4 })
	fmt.Println(fourthAge)
	createTransformFunc := createTransformFunc(5)
	sixthAge := getTransformVal(ages, createTransformFunc)
	fmt.Println(sixthAge)
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

// closure method
func createTransformFunc(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
