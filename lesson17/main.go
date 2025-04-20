package main

import (
	"fmt"
	"math"
)

func test(someFunction func(int) int, x int) {
	fmt.Println(someFunction(x))
}

func test2(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	a := func() {
		fmt.Println("hello world")
	}

	a()

	gipotenuza := func(k1 float64, k2 float64) float64 {
		return math.Sqrt(k1*k1 + k2*k2)
	}

	fmt.Println(gipotenuza(3, 4))

	square := func(x int) int {
		return x * x
	}
	test(square, 225)
	test2("hello") // not working!!
	test2("hello")()
}
