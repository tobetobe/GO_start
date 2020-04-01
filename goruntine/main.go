package main

import (
	"fmt"
	"time"
)

func send(ch chan int, i int) {
	time.Sleep(time.Second)
	ch <- i
}
func run(ch chan int) {
	for i := 0; i < 10; i++ {
		go send(ch, i)
	}
}
func sayHello() {
	fmt.Println("Hello")
}
func main() {
	ch := make(chan int, 10)
	run(ch)
	sayHello()
	var list []int
	for {
		receive, ok := <-ch
		fmt.Println(ok)
		list = append(list, receive)
		fmt.Println(list)
		if len(list) == 10 {
			break
		}
	}

}
