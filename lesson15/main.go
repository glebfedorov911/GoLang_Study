package main

import "fmt"

func main() {
	var money map[string]int = map[string]int{
		"dollars": 1000,
		"rubles":  1000000,
		"euros":   10,
	}
	fmt.Println(money)
	fmt.Println(money["dollars"])

	money["rubles"] = 10000
	fmt.Println(money)
	fmt.Println(money["rubles"])

	delete(money, "rubles")
	fmt.Println(money)
	fmt.Println(money["rubles"])
	/* несуществующее значение возвращает
	дефолтное значение у int - 0
	*/

	el, status := money["dollars"]
	// el, status := money["rubles"]
	if status {
		fmt.Println(status, el)
	} else {
		fmt.Println(status, el)
	}
}
