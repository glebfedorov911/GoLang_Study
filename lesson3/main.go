package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите члены квадратного уравнения a, b, c")

	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	var D float64
	D = (b*b - 4*a*c)
	if D > 0 {
		var x1, x2 float64
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("Найдено 2 корня:", x1, x2)
	} else if D == 0 {
		var x float64
		x = (-b) / (2 * a)
		fmt.Println("Найден 1 корень:", x)
	} else {
		fmt.Println("Корни не найдены, D < 0")
	}
}
