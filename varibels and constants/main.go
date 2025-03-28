package main

import "fmt"

func main() {
	//1 method
	var name string = "chamod"
	var age int = 22
	var isMale bool = true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMale)

	//2 method using := sign just
	fName := "udara"
	fmt.Print(fName)

	var celcius float32 = 10.32
	fmt.Print(celcius)

	var firstName, lastName string = "chamod", "udara"
	fmt.Println(firstName)
	fmt.Println(lastName)

	//print type use  printf
	fmt.Printf("%T", celcius)

	const PI float32 = 3.14
	fmt.Print(PI)

	age = 20
	name = "dulit"

	fmt.Println(name)
	fmt.Println(age)
}

//Println - add to new line
