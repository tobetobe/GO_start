package main 
import "fmt"


// func log(array [...]x) {
// 	for index,item := range x {
// 		fmt.Println(index, item)
// 	}
// }


func main() {
	array1 := [5]string{"a", "b", "c", "d", "e"}

	// fmt.Println(array1)
	// fmt.Println(&array1)

	slice1 := array1[:]

	fmt.Println(slice1)

	fmt.Println(slice1[1:1], len(slice1[1:1]), slice1[1:2], len(slice1[1:2]))
}