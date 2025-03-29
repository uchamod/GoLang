package main

import (
	"errors"
	"fmt"
)

func sum(allValue, newValues int) int {
	allValue = allValue + newValues
	return allValue
}

// func callback(func(int,int)int,int)int{

// }

func adding() {
	allValue := 100
	newValues := 25

	sum(allValue, newValues)
	fmt.Println(allValue)
}

// return multiple values
func getPoint() (int, int) {
	return 3, 4
}

// early return and error handling
func divide(divident, divisor int, err error) (int, error) {
	if divident == 0 {
		return 0, errors.New("cannot divide by zero") //return with error 
	}
	return divisor / divident, nil
	//nil means no error
}

func getCoodes() (x, y int) {
	return //automatically return x and y
	//int initial vaalues sre zero
}
func main() {
	//in go only values update when return the value
	//use = for updating
	allValue := 100
	newValues := 25

	allValue = sum(allValue, newValues)
	fmt.Println(allValue)
	//val := sum(10, 20)
	//fmt.Println(val)

	//ignore y value using _
	x, _ := getPoint()
	fmt.Println(x)

}
