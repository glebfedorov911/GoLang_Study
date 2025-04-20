package main

import "fmt"

type User struct {
	name     string
	age      int64
	password string
}

func changeUsername(user *User, newUsername string) {
	user.name = newUsername
}

func main() {
	//1 cпособ
	var user User = User{name: "John", age: 23, password: "password"}
	fmt.Println(user)

	//2 способ
	// user2 := User{name: "John", age: 23, password: "password"}//порядок не важен
	user2 := User{"John", 23, "password"} //порядок важен
	fmt.Println(user2)

	fmt.Println(user2.name)
	fmt.Println(user2.age)
	fmt.Println(user2.password)

	user2.age = 24
	fmt.Println(user2)

	changeUsername(&user2, "Gleb")
	fmt.Println(user2)
}
