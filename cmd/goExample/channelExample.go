package main

import (
	"fmt"
)

func sum(arr []int, c chan int)  {
	total := 0
	for _, v := range arr{
		total += v
		fmt.Printf("value is =>%d\n",v)
	}
	c <- total
}

//func main()  {
//	a := []int{1, 2, 3, -5, 4, 0}
//	c := make(chan int)
//
//	go sum(a[:len(a)/2], c)
//	go sum(a[len(a)/2:], c)
//
//	x := <-c
//	y := <-c
//	fmt.Println(x, y, x+y)
//	//for ; ;  {
//	//	time.Sleep(1)
//	//}
//}

func main()  {
	c := make(chan int, 4)  //有缓冲阻塞的channel，当value为0时，channel是无缓冲阻塞的
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)

}
