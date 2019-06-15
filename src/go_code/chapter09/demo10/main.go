package main

import (
	"fmt"
	"time"
)

func main() {

	//获取当前时间
	now := time.Now()
	fmt.Printf("now = %v now type = %T\n", now, now)

	//年月日，时分秒
	fmt.Printf("year = %v\n", now.Year())
	fmt.Printf("month = %v\n", now.Month())
	fmt.Printf("month = %v\n", int(now.Month()))
	fmt.Printf("day = %v\n", now.Day())
	fmt.Printf("Hour = %v\n", now.Hour())
	fmt.Printf("Minute = %v\n", now.Minute())
	fmt.Printf("Second = %v\n", now.Second())

	//时间的格式化
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()

	//时间戳
	fmt.Printf("unix时间戳 = %v, unixnano时间戳 = %v\n", now.Unix(), now.UnixNano())

}
