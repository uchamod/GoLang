package main

import "fmt"

type student struct {
	id   int
	name string
}

func (student) read() string {
	return "students read books"
}

// infinity loop with limitaions
func maxMassage(limit int) float64 {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 2.0 * float64(i)
		if i >= limit {
			return totalCost
		}
	}
}

// while loop : this is achive also a using for keyword
func ageChecker(initialAge int) {

	for initialAge <=  17 {
		fmt.Println("you cannot enter")
		initialAge++
		if initialAge >= 18 {
			fmt.Println("you can enter")
			return
		}
		continue
	}
}
func main() {
	s := student{
		id:   001,
		name: "chamod",
	}
	for i := 0; i < 5; i++ {
		fmt.Println(s.read())
	}
	fmt.Println(maxMassage(10))
	ageChecker(10)
}
