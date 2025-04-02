package main

import "fmt"

func main() {
	var p *int
	myName := "chmaod"
	age := 23
	myNamePointer := &myName
	p = &age

	fmt.Println(myNamePointer)
	fmt.Println(p)

	myString := "hello world"
	myStringPtr := &myString
	fmt.Println(myString)
	*myStringPtr = "hello chamod"
	fmt.Println(myString)

}
