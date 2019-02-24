package main

import (
	"fmt"
	"time"

)

// goroutine1 函数，循环内部每隔一秒钟打印一次总共执行次数
func goroutine1() {
	var idx = 0
	for {
		idx++
		fmt.Println("goroutine 1，执行第", idx, "次")

		// 睡眠等待1秒
		time.Sleep(time.Second * 1)
	}
}

// goroutine2 函数，循环内部每隔一秒钟打印一次总共执行次数
func goroutine2() {
	var idx = 0
	for {
		idx++
		fmt.Println("goroutine 2，执行第", idx, "次")

		// 睡眠等待1秒
		time.Sleep(time.Second * 1)
	}
}

func main() {
	fmt.Println("main function")
	//runtime.GOMAXPROCS(8)
	// 开启并发执行
	go goroutine1()
	go goroutine2()
	//for ; ;  {
	//	time.Sleep(1)
	//}
}
