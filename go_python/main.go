package main

import (
	"fmt"

	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	m := python.PyImport_ImportModule("sys")
	if m == nil {
		fmt.Println("import error")
		return
	}
	path := m.GetAttrString("path")
	if path == nil {
		fmt.Println("get path error")
		return
	}
	//加入当前目录，空串表示当前目录
	currentDir := python.PyString_FromString("")
	python.PyList_Insert(path, 0, currentDir)

	// size := python.PyList_GET_SIZE(path)
	// for i := 0; i < size; i++ {
	// 	item := python.PyList_GET_ITEM(path, i)
	// 	s := python.PyString_AsString(item)
	// 	fmt.Println(s)
	// }
	m = python.PyImport_ImportModule("fib")
	if m == nil {
		fmt.Println("import error")
		return
	}
	fib := m.GetAttrString("fib")
	if fib == nil {
		fmt.Println("get fib error")
		return
	}
	out := fib.CallFunction(python.PyInt_FromLong(10))
	if out == nil {
		fmt.Println("call fib error")
		return
	}
	fmt.Printf("fib(%d)=%d\n", 10, python.PyInt_AsLong(out))
}
