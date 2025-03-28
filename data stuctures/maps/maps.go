package main

import "fmt"

func main() {

	students := map[string]int{
		"chamod": 10,
		"udara":  20,
		"amal":   30,
	}
	fmt.Println(students)
	students["chamod"] = 25
	fmt.Println(students)

	teachers := map[string]int{}

	teachers["bimal"] = 23
	teachers["naml"] = 24

	fmt.Println(teachers)

	delete(students, "amal")
	fmt.Println(students)
}
