package main

import "fmt"

// define stucture for the student
type Student struct {
	fName string
	lName string
	age   int
	grade int
}

func main() {

	chamod := Student{
		fName: "chamod",
		lName: "udara",
		age:   23,
		grade: 12,
	}

	fmt.Println(chamod)

	fmt.Printf("full details of the  student is %+v", chamod)
	fmt.Printf("full details of the  student is %v", chamod.age )

}
