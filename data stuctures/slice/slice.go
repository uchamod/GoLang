package main

import "fmt"

func main() {

	//slice is an flexible array

	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)

	numbers = append(numbers, 5, 6, 7)
	fmt.Println(numbers)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 9} 
	fmt.Println(nums)

	nums = append(nums[:4], nums[5:]...)//use elips
	fmt.Println(nums)


}
