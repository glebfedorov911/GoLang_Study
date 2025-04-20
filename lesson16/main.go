package main

import "fmt"

func first() {
	fmt.Println("Hello from my first function")
}

func sum(a int, b int) {
	fmt.Println(a + b)
}

func sum2(a int, b int) int {
	return a + b
}

func math_func(a int, b int) (int, int, int, float64) {
	sum := a + b
	sub := a - b
	mul := a * b
	div := float64(a) / float64(b)

	return sum, sub, mul, div
}

func math_func2(a int, b int) (sum int, sub int, mul int, div float64) {
	sum = a + b
	sub = a - b
	mul = a * b
	div = float64(a) / float64(b)

	return
}

func main() {
	for i := 0; i < 10; i++ {
		first()
	}
	sum(10, 28)
	a := sum2(25, 31)
	fmt.Println(a)
	sm, sb, ml, dv := math_func(25, 6)
	fmt.Println(sm, sb, ml, dv)

	fmt.Println(math_func2(5, 2))
}
