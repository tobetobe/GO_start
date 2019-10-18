package main

import "fmt"

func main() {
	// fmt.Printf("hello, world\n")

	// var v2 = 2

	// var v1 string = "abc"

	// v3 := 10

	// fmt.Print(v1, v2, v3)

	// v3 = 500

	// fmt.Print(v3)

	// fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
	// fmt.Print("Καλημέρα κόσμε; or こんにちは 世界\n")
	// fmt.Println("Καλημέρα κόσμε; or こんにちは 世界")

	// fmt.Print(1, 2)
	// fmt.Printf(1, 2)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(
		Sunday,
		Monday,
		Tuesday,
		Wednesday,
		Thursday,
		Friday,
		Saturday)

}
