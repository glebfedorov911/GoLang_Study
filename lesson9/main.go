package main

import (
	"fmt"
	"math"
)

func main() {
	// объявление массива 1
	var names [3]string

	names[0] = "Jordan"
	fmt.Println(names[0])

	names[1] = "John"
	names[2] = "Gleb"
	fmt.Println(names)

	// объявление массива 2
	names2 := [3]string{"Gleb", "Jordan", "John"}
	fmt.Println(names2)

	fmt.Println(len(names2))

	for idx, name := range names {
		fmt.Println(idx, name)
	}

	for _, name := range names2 {
		fmt.Println(name)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	marks := [5]float64{5, 4, 3, 4, 2}
	var sum, count, middle_arithmetics float64

	for _, mark := range marks {
		sum += mark
		count += 1
	}

	middle_arithmetics = sum / count
	fmt.Println(middle_arithmetics)
	fmt.Println("Итоговая оценка:", math.Round(middle_arithmetics))

}
