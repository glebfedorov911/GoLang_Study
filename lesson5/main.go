package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Hello %d\n", i)
	}

	for i := 1; i <= 50; i++ {
		if i%37 == 0 {
			break
		}
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for idx, elem := range nums {
		fmt.Println("index:", idx, "element value:", elem)
	}

	for _, elem := range nums {
		fmt.Println(elem)
	}

	matrix := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	for _, row := range matrix {
		for _, elem := range row {
			fmt.Printf("%d\t", elem)
		}
		fmt.Printf("\n")
	}

	// for {
	// 	fmt.Println("Hello")
	// 	//бесконечный цикл
	// }
}
