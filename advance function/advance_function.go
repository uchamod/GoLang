package main

import (
	"fmt"
)

// heres talk function outside the onther function and anonymous function
// take func as parameters
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

// advance go function
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func getLogger(fomatter func(string, string) string) func(string, string) {
	return func(s1, s2 string) {
		fmt.Println(fomatter(s1, s2))
	}
}
func fomatter(s1, s2 string) string {
	return s1 + " " + s2
}

// take 3 ints and function as parameters
func calculation(a, b, c int, arithmatic func(x, y int) int) int {
	return arithmatic(arithmatic(a, b), c)
}

// anonymous function
func doMath(f func(x int) int, numbers []int) []int {
	var result []int
	for _, n := range numbers {
		result = append(result, f(n))
	}
	return result
}
func main() {
	fmt.Println(calculation(5, 6, 7, add))
	fmt.Println(calculation(5, 6, 7, sub))

	doublefunc := selfMath(add)
	fmt.Println(doublefunc(5))

	mergeWord := getLogger(fomatter)

	mergeWord("chmaod", "udara")

	//anonymous one
	nums := []int{1, 2, 3, 4, 5}
	resultData := doMath(func(x int) int {
		return x + x
	}, nums) 

	fmt.Println(resultData)
}
