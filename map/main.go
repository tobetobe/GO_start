package main

import (
	"fmt"

	goid "github.com/JakeHL/Goid"
)

type s struct {
	m map[string]*goid.UUID
}

func main() {
	s := &s{}
	fmt.Println(s.m)
	s.m = make(map[string]*goid.UUID)
	fmt.Println(s.m["123"] == nil)
}
