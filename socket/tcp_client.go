package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}
	input := bufio.NewReader(os.Stdin)

	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed, err:%v\n", err)
			return
		}
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed, err:%v\n", err)
			return
		}
		fmt.Println("received data from server: ", string(buf[:n]))
	}
}
