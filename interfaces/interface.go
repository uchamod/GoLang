package main

import (
	"fmt"
	"time"
)

// interface
type massage interface {
	getMassage() string
}

// struct
type wishMassage struct {
	birthTime      time.Time
	wishingMassage string
}
type blameMassage struct {
	blameTime    time.Time
	blameMassage string
}

func (wm wishMassage) getMassage() string {
	return fmt.Sprintf("hi %v", wm)
}

func (bm blameMassage) getMassage() string {
	return fmt.Sprintf("hi %v", bm)
}

func main() {
	   
	wm := wishMassage{ 
		birthTime:     time.Date(),
		wishingMassage: "happy birth day",
	}

	fmt.Println(wm.getMassage())
}
