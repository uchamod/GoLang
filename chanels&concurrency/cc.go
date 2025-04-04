package main

import (
	"fmt"
	"time"
)

// concurrent function
func sendEmail(massage string) {
	go func() {
		time.Sleep(time.Microsecond * 250)
		fmt.Printf("email recived : %s", massage)
		fmt.Println()
	}()
	fmt.Printf("email send,%s", massage)
	fmt.Println()

}

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}
func main() {
	test("hello how are you dear?")
	test("are you kidding me")
	test("nice try")

}
