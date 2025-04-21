package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("lesson24/txt2.txt")
	if err != nil {
		fmt.Println("ОШИБКА -", err)
		os.Exit(1)
	}

	defer file.Close()

	data := "Hello\nFrom\nGo\n\tTab"
	file.WriteString(data)

	file_data, _ := os.ReadFile(file.Name())
	fmt.Println(file_data)         //bytes slice
	fmt.Println(string(file_data)) //string

	// os.Remove("lesson24/txt2.txt")
}
