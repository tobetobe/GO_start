package main

import (
	"fmt"
	"time"
)

// func send(ch chan int, i int) {
// 	time.Sleep(time.Second)
// 	ch <- i
// }
// func run(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		go send(ch, i)
// 	}
// }
// func sayHello() {
// 	fmt.Println("Hello")
// }
// func main() {
// 	ch := make(chan int)
// 	send(ch, 2)
// 	<-ch
// 	fmt.Println("123")

// }

func main() {
	go func() {
		time.Sleep(time.Second)
		fmt.Println("hh")
	}()
	// time.Sleep(2 * time.Second)
}
