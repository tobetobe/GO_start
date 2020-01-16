package main

import (
	"bufio"
	"fmt"
	"net"

)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("received data: ", recv)
		conn.Write([]byte("ok"))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err: %v\n", err)
			continue
		}
		go process(conn)
	}
}
