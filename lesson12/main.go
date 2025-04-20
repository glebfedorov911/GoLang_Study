package main

import "fmt"

func main() {
	slice := []int{1, 4, 5, 2, 6, 9, 10}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for idx, elem := range slice {
		fmt.Printf("%d: %d\n", idx, elem)
	}
}
