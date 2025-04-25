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

// package main

// import "fmt"

// func main() {
// 	var a int
// 	fmt.Scan(&a)

// 	n1, n2, n3, n4, n5, n6 := a/100000, a/10000%10, a/1000%10, a/100%10, a/10%10, a%10
// 	if n1+n2+n3 == n4+n5+n6 {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var n, sum int

// 	fmt.Scan(&n)
// 	for i := 0; i < n; i++ {
// 		var v int
// 		fmt.Scan(&v)
// 		if 10 <= v && v <= 99 && v%8 == 0 {
// 			sum += v
// 		}
// 	}
// 	fmt.Println(sum)
// }

// package main

// import "fmt"

// func main() {
// 	var n int
// 	for {
// 		fmt.Scan(&n)
// 		if n < 10 {
// 			continue
// 		} else if n > 100 {
// 			break
// 		}
// 		fmt.Println(n)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var x, p, y, years int
// 	fmt.Scan(&x, &p, &y)

// 	for x < y {
// 		x += int(float64(x) / 100 * float64(p))
// 		years++
// 	}
// 	fmt.Println(years)
// }

// package main

// import "fmt"

// func findAllNumsInNum(num int) *[]int {
// 	var array []int
// 	delimiter := 1
// 	for num > delimiter {
// 		findNum := num / delimiter % 10
// 		array = append(array, findNum)
// 		delimiter *= 10
// 	}
// 	reverseArray(&array)
// 	return &array
// }

// func reverseArray(array *[]int) {
// 	for i, j := 0, len(*array)-1; i < j; i, j = i+1, j-1 {
// 		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
// 	}
// }

// func main() {
// 	var n1, n2 int
// 	fmt.Scan(&n1, &n2)
// 	numsInN1 := *findAllNumsInNum(n1)
// 	numsInN2 := *findAllNumsInNum(n2)
// 	for i := 0; i < len(numsInN1); i++ {
// 		for j := 0; j < len(numsInN2); j++ {
// 			if numsInN1[i] == numsInN2[j] {
// 				fmt.Print(numsInN2[j], " ")
// 			}
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var n float64
// 	fmt.Scan(&n)
// 	if n <= 0 {
// 		fmt.Printf("число %4.2f не подходит", n)
// 	} else if n > 10000 {
// 		fmt.Printf("%e", n)
// 	} else {
// 		fmt.Printf("%.4f", n*n)
// 	}
// }

package main

import "fmt"

func main() {
	fmt.Println(`
	fdsfdfsdfsd
	`)
	fmt.Println("fkksdksdfk")
}
