package main

import "fmt"

func main() {

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Print(numbers)
	fmt.Println(numbers[0])

	var names [3]string
	names[0] = "amal"
	names[2] = "kml"
	fmt.Println(names)

	var twoArray = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(twoArray)

}
