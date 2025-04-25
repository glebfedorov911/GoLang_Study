// package main

// import "fmt"

// func main() {
// 	// 	const (
// 	// 		Mon = iota
// 	// 		Tue
// 	// 		Wen
// 	// 		Thr
// 	// 		Frd
// 	// 		Sat
// 	// 		Sun
// 	// 	)
// 	// 	fmt.Println(Frd)
// 	v := 42
// 	switch v {
// 	case 100:
// 		fmt.Println(100)
// 		fallthrough
// 	case 42:
// 		fmt.Println(42)
// 		fallthrough
// 	case 1:
// 		fmt.Println(1)
// 		fallthrough
// 	default:
// 		fmt.Println("default")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var inputData int
// 	fmt.Scan(&inputData)

// 	n1, n2, n3 := inputData/100, inputData%100/10, inputData%10
// 	fmt.Println(n1, n2, n3)
// 	if n1 == n2 || n2 == n3 || n1 == n3 {
// 		fmt.Println("NO")
// 	} else {
// 		fmt.Println("YES")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	const MAX_VALUE int = 100000
// 	var a int
// 	delimiter := 1

// 	fmt.Scan(&a)
// 	for delimiter <= MAX_VALUE {
// 		if a < delimiter {
// 			fmt.Println(a / (delimiter / 10))
// 			break
// 		}
// 		delimiter *= 10
// 	}
// }

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	n1, n2, n3, n4, n5, n6 := a/100000, a/10000%10, a/1000%10, a/100%10, a/10%10, a%10
	if n1+n2+n3 == n4+n5+n6 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
