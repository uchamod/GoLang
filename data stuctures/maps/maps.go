package main

import (
	"errors"
	"fmt"
)

//slice merger

func getUserMap(names []string, numbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(numbers) {
		return nil, errors.New("slice size mismathing")
	}
	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{
			name:        names[i],
			phoneNumber: numbers[i],
		}
	}
	return userMap, nil

}

type user struct {
	name        string
	phoneNumber int
}

// delete user
type members struct {
	name                string
	age                 int
	isSheduledForDelete bool
}

func deleteSheduledUser(membersMap map[string]members, name string) (isDeleted bool, err error) {
	member, exist := membersMap[name]
	if !exist {
		return false, errors.New("user not exist")
	}
	if exist && !member.isSheduledForDelete {
		return false, nil
	}
	delete(membersMap, name)
	return true, nil
}

func getCount(userIds []string) map[string]int {
	if len(userIds) == 0 {
		return nil
	}
	countMap := make(map[string]int)
	for i := 0; i < len(userIds); i++ {
		s := userIds[i]
		user, exist := countMap[s]
		if exist {
			user++
			countMap[s] = user
		} else {
			countMap[s] = 1
		}
	}
	return countMap
}

func main() {

	students := map[string]int{
		"chamod": 10,
		"udara":  20,
		"amal":   30,
	}
	fmt.Println(students)
	students["chamod"] = 25
	fmt.Println(students)

	teachers := map[string]int{}

	teachers["bimal"] = 23
	teachers["naml"] = 24

	fmt.Println(teachers)

	delete(students, "amal")
	fmt.Println(students)

	fmt.Println(getUserMap([]string{"amal", "kmal", "vimal"}, []int{123, 456, 789}))

	memberMap := make(map[string]members)
	memberMap["chamod"] = members{name: "chamod", age: 22, isSheduledForDelete: false}
	memberMap["amal"] = members{name: "amal", age: 22, isSheduledForDelete: false}
	memberMap["kmal"] = members{name: "kmal", age: 23, isSheduledForDelete: true}

	fmt.Println(deleteSheduledUser(memberMap, "saman"))
	userIds := []string{"123", "456", "789", "456", "123", "967", "123"}
	fmt.Println(getCount(userIds))
}
