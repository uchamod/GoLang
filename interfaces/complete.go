package main

import (
	"fmt"
)

// switching interface
func getInterfaceReport(e expense) (string, float64) {
	switch s := e.(type) {
	case email:
		return s.body, s.cost()
	case call:
		return s.voice, s.cost()
	default:
		return "", 0.0
	}
}

type email struct {
	isSubscribed bool
	body         string
}
type call struct {
	isEnable bool
	voice    string
}

// email method
func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (e email) print() {
	fmt.Println(e.body)
}

// call method
func (c call) cost() float64 {
	if c.isEnable {
		return float64(len(c.voice)) * 0.8
	}
	return float64(len(c.voice)) * 0.1
}
func (c call) print() {
	fmt.Println(c.voice)
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

func print(p printer) {
	p.print()
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test(e, e)
  
	c := call{
		isEnable: false,
		voice:    "hi there",
	}

	fmt.Println(getInterfaceReport(e))
	fmt.Println(getInterfaceReport(c)) 
}
