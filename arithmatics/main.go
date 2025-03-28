package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 int = 20
	var sum int = num1 + num2

	fmt.Println(sum)

	findWithCircle()

}

func findWithCircle() {
	fmt.Println("calculate circul area or ciriculam")

	fmt.Println("1.Area")
	fmt.Println("2.ciriculam")
	var radius float32
	var choice int
	fmt.Println("Enter your preference:")
	fmt.Scan(&choice)
	fmt.Println("Enter radius:")
	fmt.Scan(&radius)
	calculate(radius, choice)
}

func calculate(radius float32, choice int) {
	var result float32
	if choice == 1 {
		result = 3.14 * radius * radius
		fmt.Println(result)
		return
	} else {
		result = 2 * 3.14 * radius
		fmt.Println(result)
		return
	}
}
