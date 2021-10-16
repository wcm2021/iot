package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("listen error:", err)
				break
			}

			go handleConn(conn)
		}
	}()

	//client
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		n, err := conn.Write([]byte("dddd"))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("send count = ", n)

		time.Sleep(time.Second)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	recvBuf := make([]byte, 100)
	for {
		n, err := conn.Read(recvBuf)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("recv buf =", string(recvBuf[:n]))
	}
}
