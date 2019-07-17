package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		user string
		pwd  string
		host string
		port int
	)

	//类似MySQL的登录语句
	//终端输入go run main.go -u root -pwd root -h 192.168.0.1 -port 3306

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")
	flag.Parse()
	fmt.Printf("user = %v, pwd = %v, host = %v, port = %v", user, pwd, host, port)

}
