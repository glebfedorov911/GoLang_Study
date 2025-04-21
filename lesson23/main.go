package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	// go say("World", channel)
	// time.Sleep(2 * time.Second)
	// fmt.Println(<-channel)

	// fmt.Println("Начало программы")
	// go timer(5, channel)
	// fmt.Println(<-channel)
	// fmt.Println("Конец программы")
	// /*Если закомментировать строку 17, то программа завершится (тк основной поток стоп, в другом выполняются действия)
	// Если же не комментировать строку 17, а ждать полного выполнения с выводом данных из канала, то программа в основном потоке*/

	go myfunc(channel)
	for elem := range channel {
		fmt.Println(elem)
	}

	fmt.Println(<-channel) //пустой канал

}

func say(greet string, channel chan int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Hello %s!!!\n", greet)
	channel <- 7
}

func timer(seconds int, channel chan int) {
	time.Sleep(time.Duration(seconds) * time.Second)
	channel <- seconds
	fmt.Println(seconds)
}

func myfunc(channel chan int) {
	for i := 0; i <= 5; i++ {
		channel <- i
	}
	close(channel)
}
