package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 5

	sum := a + b
	// + - * / ++ -- *= += -= /= %
	// / - для float - число с плавающей точкой
	// / - для int - целочисленное деление
	fmt.Println(sum)

	fmt.Println(a / b)
	var a1 float64 = 3
	var b1 float64 = 5
	fmt.Println(a1 / b1)

	var num float64
	var result float64

	num = 144
	result = math.Sqrt(num)
	fmt.Println(result)

	num = 25.273773
	result = math.Ceil(num) // к большему
	fmt.Println(result)
	result = math.Floor(num) // к меньшему
	fmt.Println(result)
	result = math.Round(num) // по правилам математики
	fmt.Println(result)
	result = math.Round(25.5) // по правилам математики
	fmt.Println(result)

	var symbol string
	var r, n1, n2 float64

	for {
		fmt.Println("Какое действие вы хотите выполнить (*, +, -, /)")
		fmt.Scan(&symbol)

		fmt.Println("Введите числа для выполнения операции")
		fmt.Scan(&n1, &n2)

		switch symbol {
		case "+":
			r = n1 + n2
		case "-":
			r = n1 - n2
		case "*":
			r = n1 * n2
		case "/":
			r = n1 / n2
		default:
			fmt.Println("Вы ввели недопустимый символ")
		}

		fmt.Println("Результат:", r)
	}
}
