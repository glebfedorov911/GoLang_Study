package main

import "fmt"

func main() {
	const DEGREE_HOURS = 30
	var d int
	fmt.Scan(&d)
	hours := d / DEGREE_HOURS
	minutes := (d - hours*DEGREE_HOURS) * 2
	fmt.Println("It is", hours, "hours", minutes, "minutes")
}
