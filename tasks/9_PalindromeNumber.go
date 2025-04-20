package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	stringX := strconv.Itoa(x)
	reverse_string := []rune(stringX)

	for i, j := 0, len(reverse_string)-1; i < j; i, j = i+1, j-1 {
		reverse_string[i], reverse_string[j] = reverse_string[j], reverse_string[i]
	}
	return string(reverse_string) == stringX
}
