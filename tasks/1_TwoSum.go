package main

func twoSum(nums []int, target int) []int {
	numsObratniy := map[int]int{}
	var result []int
	for idx, value := range nums {
		numsObratniy[target-value] = idx
	}
	for idx, value := range nums {
		el, status := numsObratniy[value]
		if status && idx != el {
			result = append(result, el)
			result = append(result, idx)
			break
		}
	}
	return result
}
