package main

import (
	"errors"
	"fmt"
)

// this will take any type of thing as parameter it can be slice ,array etc
// and it will return same type of things will enter
func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func getLast[T any](s []T) T {

	length := len(s)
	if length == 0 {
		//difine empty something (email or payment struct)
		var zero T
		return zero
	}
	last := s[length-1]

	return last
}

type email struct {
	massage        string
	senderEmail    string
	reciptionEmail string
}

type payment struct {
	amount int

	senderEmail string
}

// payment charging fuunction
func chargForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	if newItem.GetCost() > balance {
		return oldItems, balance, errors.New("insufficient balnce")
	}
	newItemSlice := append(oldItems, newItem)
	balance = balance - newItem.GetCost()
	return newItemSlice, balance, nil
}

type lineItem interface {
	GetCost() float64
	GetName() string
}
type subscription struct {
	name string
	cost float64
}
type oneTimePayment struct {
	name string
	cost float64
}

func (s subscription) GetCost() float64 {
	return s.cost
}
func (o oneTimePayment) GetCost() float64 {
	return o.cost
}
func (s subscription) GetName() string {
	return s.name
}
func (o oneTimePayment) GetName() string {
	return o.name
}

func main() {
	fSlice, sSlice := splitAnySlice([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(fSlice)
	fmt.Println(sSlice)
	mails := []email{{
		massage:        "hi",
		senderEmail:    "amal",
		reciptionEmail: "kmal",
	},
		{
			massage:        "hey",
			senderEmail:    "namal",
			reciptionEmail: "ranil",
		},
	}

	payments := []payment{{
		amount:      2,
		senderEmail: "amal",
	},
		{
			amount:      1,
			senderEmail: "ranil",
		},
	}
	//charge func
	newItem := subscription{
		name: "subscription",
		cost: 25.7,
	}

	oldItems := []subscription{
		{
			name: "subscription",
			cost: 20.0,
		},
		{
			name: "one time payment",
			cost: 10.6,
		},
		{
			name: "one time payment",
			cost: 12.6,
		},
	}

	fmt.Println(getLast(mails))
	fmt.Println(getLast(payments))
	fmt.Println(chargForLineItem(newItem, oldItems, 30.0))
}
