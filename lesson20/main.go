package main

import "fmt"

type User struct {
	name     string
	age      int64
	password string
	score    []int
}

func (user User) getName() string {
	return user.name
}

func (user *User) setName(username string) {
	user.name = username
}

func (user User) checkAgeToValid() bool {
	var VALID_USER_AGE_TO_ACCESS int64 = 18
	return user.age >= VALID_USER_AGE_TO_ACCESS
}

func (user User) getMaxScores() int {
	max_idx := 0
	for idx, elem := range user.score {
		if elem > user.score[max_idx] {
			max_idx = idx
		}
	}
	return user.score[max_idx]
}

func main() {
	user := User{"John", 25, "pass", []int{23, 67, 89, 102, 345, 1}}
	fmt.Println(user)

	name := user.getName()
	fmt.Println(name)

	user.setName("Gleb")
	fmt.Println(user)

	is_valid := user.checkAgeToValid()
	fmt.Println((is_valid))

	fmt.Println(user.getMaxScores())
}
