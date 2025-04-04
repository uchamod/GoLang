package main

import (
	"fmt"
	"time"
)

func filterOldEmails(emails []Email) {
	isOldMail := make(chan bool)
	go func() {
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldMail <- true //send  data to chanel
				//this must wait to get onther vallue until current value recived by onthr variable
				continue
			}
			isOldMail <- false
		}
	}()

	isOld := <-isOldMail //recive data from chanel
	fmt.Println("1 email is old", isOld)
	isOld = <-isOldMail
	fmt.Println("2 email is old", isOld)
	isOld = <-isOldMail
	fmt.Println("3 email is old", isOld)
}

type Email struct {
	body string
	date time.Time
}

func test(emails []Email) {
	filterOldEmails(emails)
	fmt.Println("---")
}
func main() {
	test([]Email{
		{
			body: "hi how are you?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "hi hows  you?",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "hinjdiu?",
			date: time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})

}
