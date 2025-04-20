package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

type NumbersInterface interface {
	Sum() int
	Multiply() int
	Substract() int
	Division() float64
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}

func (n Numbers) Multiply() int {
	return n.num1 * n.num2
}

func (n Numbers) Substract() int {
	return n.num1 - n.num2
}

func (n Numbers) Division() float64 {
	return float64(n.num1) / float64(n.num2)
}

func main() {
	var numbersInterface NumbersInterface
	numbers := Numbers{9, 8}
	numbersInterface = numbers
	fmt.Printf("Sum %d\n", numbersInterface.Sum())
	fmt.Printf("Multiply %d\n", numbersInterface.Multiply())
	fmt.Printf("Division %f\n", numbersInterface.Division())
	fmt.Printf("Substract %d\n", numbersInterface.Substract())
}
