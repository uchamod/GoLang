package main

import "fmt"

// define stucture for the student
type Student struct {
	fName string
	lName string
	age   int
	grade int
}
type massageToSend struct {
	phoneNumber int
	massage     string
}

//formatting string use printf and %v for any type of data type

// use struct as parameter
func test(m massageToSend) {
	fmt.Printf("sending massage : %v to mobile number %v", m.massage, m.phoneNumber)

}

//struct can hold onther struct

type Wheel struct {
	radius   float32
	material string
}

// can use any type  of datatype
type Car struct {
	model     string
	year      int
	price     int
	wheeltype Wheel
}

// inheritance struct
type truck struct {
	model string
	make  string
}

type lorray struct {
	truck
	year  int
	price int
}

// methods on struct
type rect struct {
	width  float32
	height float32
}

// in struct methods the method name is use after  the parameter
func (r rect) area() float32 { //this is belongs to 'rect' struct
	return r.height * r.width
}
func main() {
	//anonymous struct only use one point
	newCar := struct {
		model string
		make  string
	}{
		model: "ford",
		make:  "3.0",
	}

	fmt.Println(newCar)
	chamod := Student{
		fName: "chamod",
		lName: "udara",
		age:   23,
		grade: 12,
	}

	test(massageToSend{
		phoneNumber: 342927363,
		massage:     "test massage",
	})
	fmt.Println(chamod)

	fmt.Printf("full details of the  student is %+v", chamod)
	fmt.Printf("full details of the  student is %v", chamod.age)

	//inheritance struct
	lalend := lorray{
		truck: truck{ //key is same as varible
			model: "2.0",
			make:  "ashoke",
		},
		year:  2025,
		price: 1000000,
	}

	fmt.Println(lalend.truck)

	r := rect{
		width:  12.5,
		height: 10.8,
	}

	fmt.Println(r.area())
}
