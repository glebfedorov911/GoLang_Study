package main

import (
	"fmt"
	"os"
)

func readFile(title string) string {
	file_data, err := os.ReadFile(title)

	if err == nil {
		return string(file_data)
	}
	return "Cannot read file"
}

func writeFile(title string, description string, rights int) bool {
	data := []byte(description)
	err := os.WriteFile(title, data, os.FileMode(rights))
	return err == nil
}

func removeFile(title string) {
	os.Remove(title)
}

func main() {
	file_data := readFile("lesson22/story.txt")
	writeFile("lesson22/myFirstWriteFileOnGo.txt", file_data, 0600)
	file_data = readFile("lesson22/myFirstWriteFileOnGo.txt")
	fmt.Println(file_data)
	removeFile("lesson22/myFirstWriteFileOnGo.txt")
}
