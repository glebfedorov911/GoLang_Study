package main

import "fmt"

func main() {
	var age int8 = 23
	// int - от -беск до +беск
	// uint - от 0 до +беск
	fmt.Println(age)

	var number float64 = 275.672
	fmt.Println(number)

	seconds := 32
	fmt.Println(seconds)

	var name string = "Gleb"
	fmt.Println(name)

	job := "programmer"
	fmt.Println(job)
	fmt.Println(len(job))

	var username string
	fmt.Println("What is your name? :)")
	fmt.Scan(&username)
	fmt.Println("Your name is", username)
	fmt.Println("Your name is " + username)

	var userAge int8
	fmt.Println("How old are you?")
	fmt.Scan(&userAge)
	fmt.Println("You are " + fmt.Sprint(userAge) + " years!")

	var n1 int8 = 2
	var n2 int64 = int64(n1)

	fmt.Println(n1, n2)

	var n3 int64 = 3123
	var n4 int8 = int8(n3)

	fmt.Println(n3, n4)
}
