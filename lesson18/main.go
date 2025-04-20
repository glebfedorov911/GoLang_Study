package main

import "fmt"

func change(str string) {
	str = "random"
}

func change2(str *string) {
	*str = "random"
}

func main() {
	a := 5
	pointer := &a // адрес в памяти
	fmt.Println(pointer)

	fmt.Println(*pointer) // получаем значение из ссылки
	// a == *&a == *pointer

	*pointer = 25
	fmt.Println(a)

	s := "LLL"
	fmt.Println(s)
	change(s) // ничего не изменится
	fmt.Println(s)

	change2(&s) // изменится
	fmt.Println(s)

	var pointer2 *string = &s
	fmt.Println(pointer2, *pointer2)
}
