package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "172.16.24.203:8888")
	if err != nil {
		fmt.Println("client dial err =", err)
		return
	}
	fmt.Println("conn success, conn =", conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err =", err)
		}

		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出....")
			break
		}

		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err =", err)
		}
		fmt.Printf("客户端发送了%d字节的数据\n", n)
	}

}
