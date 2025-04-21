package main

import "fmt"

func main() {
	i := 0
LOOP:
	if i > 50 {
		goto END
	}
	fmt.Println(i)
	i++
	goto LOOP
END:
	fmt.Println("final")
}

// НЕ ИСПОЛЬЗУЙТЕ GOTO
