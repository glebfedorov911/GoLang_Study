package main

import "fmt"

func main() {
	age := 30

	fmt.Println("My age is " + fmt.Sprint(age))

	name := "John"
	money := 5015.132
	fmt.Printf("My age is %d\n", age)
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("I have %f dollars\n", money)
	/*
		%d - int (decimal)
		%s - string
		%f - float
		%t - boolean
	*/
}
