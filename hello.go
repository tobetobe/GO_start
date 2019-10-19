package main

import (
	"fmt"
	"reflect"
	"strings"
)

// var a string = "G"

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

	// const (
	// 	Sunday = iota
	// 	Monday
	// 	Tuesday
	// 	Wednesday
	// 	Thursday
	// 	Friday
	// 	Saturday
	// )

	// fmt.Println(
	// 	Sunday,
	// 	Monday,
	// 	Tuesday,
	// 	Wednesday,
	// 	Thursday,
	// 	Friday,
	// 	Saturday)

	// a := "LY"
	// println(a)
	// n()
	// m()
	// n()
	// var a int = 5
	// var b int64 = 5

	// fmt.Println(a == int(b))

	// a := 1 << 10

	// fmt.Println(reflect.TypeOf(float32(a)))

	// type myStr string

	// var str myStr = `
	// 	Good Morning!
	// 	Good Afternoon!😊
	// `

	// fmt.Println(str, len(str), str[0])

	var str string = "abcdefghijklmn"

	sl := strings.Split(str, "")

	fmt.Println(str)

	for index, val := range sl {
		fmt.Println(index, val)
	}

	fmt.Println(reflect.TypeOf(sl))
}

// func n() {
// 	print(a)
// }

// func m() {
// 	a = "O"
// 	print(a)
// 	c()
// }

// func c() {
// 	fmt.Printf(" c : %s ", a)
// }
