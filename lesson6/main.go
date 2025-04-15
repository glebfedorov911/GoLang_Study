package main

import "fmt"

func main() {
	name := "Kate"

	switch name {
	case "Jordan":
		fmt.Println("Hello, Jordan")

	case "Kate":
		fmt.Println("Hello, Kate")

	case "John":
		fmt.Println("Hello, John!!")

	default:
		fmt.Println("Not registered")
	}

	number := 10

	switch {
	case number > 1:
		fmt.Println("case1")
		fallthrough
	case number < 11:
		fmt.Println("case2")
	}

}
