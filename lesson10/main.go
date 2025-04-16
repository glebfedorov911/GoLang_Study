package main

import "fmt"

func main() {
	array := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(array)
	fmt.Println(array[2][1])
}
