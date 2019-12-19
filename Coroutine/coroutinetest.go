package main

import "fmt"

func Count(ch chan int) {
	fmt.Println("Counting", ch)
	ch <- 1
}

func main() {
	chs := make([]chan int, 10)

	fmt.Println("pre", chs)
	for i := 0; i < 10; i++ {
		fmt.Println(chs[i])
		chs[i] = make(chan int)
		fmt.Println("chs[i]", chs[i])
		go Count(chs[i])
	}
	for _, ch := range chs {
		// fmt.Println(ch, &ch)
		<-ch
	}
	fmt.Println("after", chs)
}
