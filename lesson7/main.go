package main

import "fmt"

func main() {
	/*
		>
		>=
		<
		<=
		==
		!=
		&& - и
		|| -  или
		! - не
	*/

	a := 1

	if a > 0 {
		fmt.Println("a>0")
	}

	b1 := 10
	b2 := 30

	if b1 > 5 && b2 > 20 {
		fmt.Println("cond1")
	}
	if b1 > 5 || b2 > 100 {
		fmt.Println("cond2")
	}
	if !(b1 > 100 && b2 > 300) {
		fmt.Println("cond3")
	}
}
