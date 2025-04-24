// package main

// import (
// 	"fmt"
// )

// func fillArray(length int, nums *[]int) {
// 	for i := 0; i < length; i++ {
// 		(*nums) = append((*nums), 0)
// 	}
// }

// func removeDuplicates(nums []int) int {
// 	busy_position := []int{}
// 	nums_copy := nums
// 	k := 0

// 	fillArray(len(nums), &busy_position)

// 	for i := 0; i < len(nums_copy); i++ {
// 		if busy_position[nums_copy[i]] == 0 {
// 			fmt.Println(k, nums[k], k-1, k+1)
// 			nums = append(nums[:k], nums[k+1:]...)
// 			k++
// 		}
// 		busy_position[nums_copy[i]] = 1
// 	}
// 	fmt.Println(nums)
// 	return k
// }

// func main() {
// 	nums := []int{1, 1, 2}
// 	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
// 	a := removeDuplicates(nums)
// 	fmt.Println(a)
// }
