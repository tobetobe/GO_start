package main

import (
	"flag"
	"fmt"
)

var (
	a string
	b string
	c int
	d int
	e bool
	f bool
)

func init() {
	flag.StringVar(&a, "a", "aaa", "usage for a")
	flag.StringVar(&b, "b", "bb", "usage for b")
	flag.IntVar(&c, "c", 1, "usage for c")
	flag.IntVar(&d, "d", 2, "usage for d")
	flag.BoolVar(&e, "e", true, "usage for e")
	flag.BoolVar(&f, "f", false, "usage for f")
}

func main() {

	flag.Parse()

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)

	others := flag.Args()
	fmt.Println(others)
}
