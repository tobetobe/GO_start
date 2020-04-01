package main

import (
	"encoding/json"
	"fmt"
)

//
type BasicMessage struct {
	SenderName string
	Recipient  string
	Message    string
}

func defaultDeserializeMessage(message string) (*BasicMessage, error) {
	b := &BasicMessage{}
	err := json.Unmarshal([]byte(message), b)
	if err != nil {
		fmt.Printf("json.Unmarshal failed err: %v\n", err)
		return b, err
	}
	// fmt.Printf("%#v\n", b)
	return b, nil
}

func main() {
	// `{"senderName":"101","recipient":"123","message":"Hello"}`
	res, err := defaultDeserializeMessage(`{"senderName":"101","recipient":"123","message":"Hello"}`)
	fmt.Printf("%#v, %#v\n", res, err)
	str := `{"Title":"101","Students":"123"}`
	c1 := &Class{}
	json.Unmarshal([]byte(str), c1)
	fmt.Printf("%#v\n", c1)
}

//Class 班级
type Class struct {
	Title    string
	Students string
}
