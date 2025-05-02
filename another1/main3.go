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

// package main

// import "fmt"

// func main() {
// 	fmt.Println(`
// 	fdsfdfsdfsd
// 	`)
// 	fmt.Println("fkksdksdfk")
// }

// package main

// import "fmt"

// func main() {
// 	var a []int
// 	fmt.Println(a, len(a), cap(a))
// 	a = append(a, 4, 3, 4)
// 	fmt.Println(a, len(a), cap(a))
// 	a = append(a, 4)
// 	fmt.Println(a, len(a), cap(a))
// 	fmt.Println(a[3:5], a)

// 	b := [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
// 	c := b[9:12]
// 	fmt.Println(len(c), cap(c))

// 	baseSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	shortSlice := baseSlice[3:5]
// 	shortSlice = append(shortSlice, 66, 77, 88)
// 	fmt.Println(baseSlice, shortSlice)
// 	shortSlice = append(shortSlice, 99, 100, 1000, 10000)
// 	fmt.Println(baseSlice, shortSlice)

// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var a, b, c float64
// 	fmt.Scan(&a, &b, &c)

// 	gipotenuza := math.Sqrt(a*a + b*b)
// 	fmt.Println(gipotenuza)
// }

// package main

// import "fmt"

// func main() {
// 	var num, v int
// 	fmt.Scan(&num)
// 	nums := make([]int, num, num)

// 	for num != 0 {
// 		fmt.Scan(&v)
// 		num--
// 		nums[num] = v
// 	}

// 	mn := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] < mn {
// 			mn = nums[i]
// 		}
// 	}

// 	var cnt int
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == mn {
// 			cnt++
// 		}
// 	}
// 	fmt.Println(cnt)
// }

// package main

// import "fmt"

// func getCifrSqrtByStep(n int) int {
// 	var cifrSqrt int
// 	for n != 0 {
// 		cifrSqrt += n % 10
// 		n /= 10
// 	}
// 	if cifrSqrt/10 == 0 {
// 		return cifrSqrt
// 	}
// 	return getCifrSqrtByStep(cifrSqrt)
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	n = getCifrSqrtByStep(n)
// 	fmt.Println(n)
// }

// package main

// import "fmt"

// func kratn7(n int) bool {
// 	part1 := n / 10
// 	part2 := n % 10

// 	return (part1-part2*2)%7 == 0
// }

// func main() {
// 	var n1, n2 int
// 	kratn := false
// 	fmt.Scan(&n1, &n2)
// 	for i := n2; i >= n1; i-- {
// 		if kratn7(i) {
// 			fmt.Println(i)
// 			kratn = true
// 			break
// 		}
// 	}
// 	if !kratn {
// 		fmt.Println("NO")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var n int
// 	var str string
// 	fmt.Scan(&n)

// 	ost := n % 10
// 	switch {
// 	case n >= 11 && n <= 20:
// 		str = fmt.Sprintf("%d korov", n)
// 	case ost == 1:
// 		str = fmt.Sprintf("%d korova", n)
// 	case ost >= 2 && ost <= 4:
// 		str = fmt.Sprintf("%d korovy", n)
// 	default:
// 		str = fmt.Sprintf("%d korov", n)
// 	}
// 	fmt.Println(str)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	max_stepen := math.Floor(math.Log2(float64(n)))
// 	cnt := -1
// 	for cnt < int(max_stepen) {
// 		cnt++
// 		fmt.Printf("%.0f ", math.Pow(2.0, float64(cnt)))
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var A int
// 	fmt.Scan(&A)

// 	i, j := 1, 1
// 	cnt := 3
// 	for i+j < A {
// 		i, j = i+j, i
// 		cnt++
// 	}
// 	if i+j == A {
// 		fmt.Println(cnt)
// 	} else {
// 		fmt.Println(-1)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	fmt.Printf("%b", n)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var n, delete_n int
// 	var nums []int
// 	fmt.Scan(&n, &delete_n)

// 	for n != 0 {
// 		ost := n % 10
// 		n /= 10
// 		if ost == delete_n {
// 			continue
// 		}
// 		nums = append(nums, ost)
// 	}

// 	var result string
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		result += fmt.Sprintf("%d", nums[i])
// 	}
// 	fmt.Println(result)
// }

// package main

// import (
// 	"fmt"
// )

// func sumInt(nums ...int) (int, int) {
// 	sum := 0
// 	for _, elem := range nums {
// 		sum += elem
// 	}
// 	return len(nums), sum
// }

// func main() {
// 	a, b := sumInt(1, 2, 3, 4)
// 	fmt.Println(a, b)
// }

// package main

// import "fmt"

// type Player struct {
// 	On    bool
// 	Ammo  int
// 	Power int
// }

// func (p *Player) Shoot() bool {
// 	if !p.On || p.Ammo == 0 {
// 		return false
// 	}
// 	p.Ammo--
// 	return true
// }

// func (p *Player) RideBike() bool {
// 	if !p.On || p.Power == 0 {
// 		return false
// 	}
// 	p.Power--
// 	return true
// }

// func main() {

// 	testStruct := Player{true, 3, 4}
// 	a := testStruct.Shoot()
// 	fmt.Println(a)
// 	/*
// 	 * Экземпляр созданной вами структуры необходимо передать в качестве
// 	 * аргумента функции testStruct, которая выполнит проверку соблюдения
// 	 * всех условий задания/
// 	 */

// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	// bs := []byte("hello world")
// 	// rs := []rune("hello world")
// 	// js := "hello world"

// 	// fmt.Println(bs[0], rs[0], js[0])
// 	// fmt.Println(string(bs[0]), string(rs[0]), string(js[0]))
// 	// fmt.Println(utf8.RuneCountInString(string(rs)))
// 	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
// 	var firstSymbol, lastSymbol string
// 	for idx, elem := range text {
// 		if idx == 0 {
// 			firstSymbol = string(elem)
// 		}
// 		if idx == len(text)-3 {
// 			lastSymbol = string(elem)
// 		}
// 	}
// 	if strings.ToUpper(firstSymbol) == firstSymbol && lastSymbol == "." {
// 		fmt.Println("Right")
// 	} else {
// 		fmt.Println("Wrong")
// 	}
// }

// package main

// import "fmt"

// func reverseString(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

// func main() {
// 	var v string
// 	fmt.Scan(&v)
// 	if v == reverseString(v) {
// 		fmt.Println("Палиндром")
// 	} else {
// 		fmt.Println("Нет")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var x, s string
// 	fmt.Scan(&x, &s)
// 	fmt.Println(strings.Index(x, s))
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var s string
// 	fmt.Scan(&s)

// 	runes := []rune(s)
// 	var r string
// 	for i := 1; i < len(runes); i += 2 {
// 		r += string(runes[i])
// 	}
// 	fmt.Println(r)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var s string
// 	fmt.Scan(&s)

// 	runes := []rune(s)
// 	var r string
// 	for i := 0; i < len(runes); i++ {
// 		if strings.Count(s, string(runes[i])) == 1 {
// 			r += string(runes[i])
// 		}
// 	}
// 	fmt.Println(r)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode/utf8"
// )

// func main() {
// 	const VOLUME_ALPHABET = "qwertyuiopasdfghjklzxcvbnm"
// 	const VOLUME_NUMS = "0123456789"

// 	var v string
// 	fmt.Scan(&v)

// 	var wasNum, wasAlphabet bool
// 	var msg string
// 	for _, elem := range []rune(v) {
// 		if strings.Contains(VOLUME_ALPHABET, strings.ToLower(string(elem))) {
// 			wasAlphabet = true
// 		} else if strings.Contains(VOLUME_NUMS, string(elem)) {
// 			wasNum = true
// 		} else {
// 			msg = "Wrong password"
// 			break
// 		}
// 	}

// 	if msg == "" && (wasNum || wasAlphabet) && utf8.RuneCountInString(v) >= 5 {
// 		msg = "Ok"
// 	} else {
// 		msg = "Wrong password"
// 	}
// 	fmt.Println(msg)
// }

// package main

// import (
// 	"fmt"
// )

// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("ошибка")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	var a, b int
// 	_, err := fmt.Scan(&a, &b)
// 	if err != nil {
// 		fmt.Println("ошибка")
// 	} else {
// 		res, err := divide(a, b)
// 		if err != nil {
// 			fmt.Println("ошибка")
// 		} else {
// 			fmt.Println(res)
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var str, result string
// 	fmt.Scan(&str)

// 	for _, elem := range []rune(str) {
// 		elemInt, _ := strconv.Atoi(string(elem))
// 		elemIntInSqrt := elemInt * elemInt
// 		result += fmt.Sprintf("%d", elemIntInSqrt)
// 	}
// 	fmt.Println(result)
// }

// package main

// import "fmt"

// func main() {
// 	v := 5
// 	p := &v
// 	fmt.Print(*p, " ")
// 	changePointer(&p)
// 	fmt.Print(*p)
// }

// func changePointer(p **int) {
// 	v := 3
// 	*p = &v
// }

// package main

// import "fmt"

// func main() {
// 	v := 5
// 	p := &v
// 	fmt.Print(*p, " ")
// 	changePointer(&v)
// 	fmt.Print(v, "|", p)
// }

// func changePointer(v *int) {
// 	*v = 3
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func T(w float64) float64 {
// 	return 6 / w
// }

// func W(k, m float64) float64 {
// 	return math.Sqrt(k / m)
// }

// func M(p, v float64) float64 {
// 	return p * v
// }

// func main() {
// 	var k, p, v float64
// 	fmt.Scan(&k, &p, &v)
// 	m := M(p, v)
// 	w := W(k, m)
// 	t := T(w)
// 	fmt.Println(t)
// }

package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   {"a1", "a2", "a3"},
		100:  {"b1", "b2", "b3"},
		1000: {"c1", "c2", "c3"},
	}
	cityPopulation := map[string]int{
		"a1": 10, "a2": 20, "a3": 30, "b1": 100, "b2": 200, "b3": 300, "c1": 1000, "c2": 2000, "c3": 3000,
	}

	for _, city := range append(groupCity[10], groupCity[1000]...) {
		delete(cityPopulation, city)
	}

	fmt.Println(cityPopulation)

	// for key, _ := range cityPopulation {
	// 	inCityPopulation := false
	// 	for _, city := range groupCity[100] {
	// 		if city == key {
	// 			inCityPopulation = true
	// 		}
	// 	}
	// 	if !inCityPopulation {
	// 		delete(cityPopulation, key)
	// 	}
	// }

	// fmt.Println(cityPopulation)
}
