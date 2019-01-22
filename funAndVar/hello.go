package main

import "fmt"

//map 是一种引用类型 一个map被赋值以后的改变会引起原先数据的修改
// map是无序的
// map的value必须使用key获得 不能使用index进行获取

func printMap()  {
	numbers := make(map[string]int)

	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Printf("未修改前：%d",numbers["one"])

	var numbers2 = numbers
	numbers2["one"] = 100

	fmt.Printf("未修改前：%d",numbers["one"])
}
//Slice 也是一种引用类型
//
func printSlice()  {
	arr := []int {9,8,7,6,5,4,3,2,1,10}

	for i := 0; i < len(arr); i++  {
		fmt.Println(arr[i])
	}
	fmt.Printf("slice的长度：%d\n", len(arr))
	fmt.Printf("slice的容量：%d\n", cap(arr))

}

func main() {
	//printMap()
	printSlice()
}
