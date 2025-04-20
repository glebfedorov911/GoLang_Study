package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 1, 2, 5, 7, 4}
	fmt.Println(slice)

	slice = append(slice, 8)
	fmt.Println(slice)

	slice[0] = 5
	fmt.Println(slice)

	sort.Ints(slice) // cортировка
	fmt.Println(slice)

	slice2 := []string{"a", "c", "d", "z", "v", "b"}
	fmt.Println(slice2)

	sort.Strings(slice2)
	fmt.Println(slice2)
}
