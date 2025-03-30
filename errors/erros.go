package main

import (
	"errors"
	"fmt"
)

func main() {
	user := User{
		id:        1,
		name:      "chamod",
		isStudent: true,
	}
	userName, err := getUser(user)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(userName)
	}
	userProfile, err := getUserProfile(user.id)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(userProfile)
	}
	newStudent := student{
		id:   1,
		name: "chamos",
	}
	fmt.Println(sendMassage("this is test massage", ""))

	newStudent.readBooks("jungle book")

}

type User struct {
	id        int
	name      string
	isStudent bool
}

func getUser(user User) (string, error) {
	if user.id != 0 {
		return user.name, nil
	} else {
		return "", errors.New("Invalid user")
	}
}
func getUserProfile(id int) (int, error) {
	if id != 0 {
		return id, nil
	} else {
		return 0, errors.New("Invalid user")
	}
}

//Error() is an interface so we can build our own error methods

type userError struct {
	name string
}
type student struct {
	id   int
	name string
}

type read interface {
	readBooks() string
}

func (student) readBooks(name string) string {
	return fmt.Sprintf("students read %s", name)
}

func getCheck() string {
	return student.readBooks(student{id: 1, name: "chamod"}, "jungle book")
}

func (e userError) Error(err string) error {
	return errors.New("this is an error")
}

func sendMassage(msg, username string) (string, error) {
	defultError := userError{
		name: "this is error",
	}
	if username == "" {
		return "", defultError.Error("this is errrrorr")
	}
	return username, nil
}
