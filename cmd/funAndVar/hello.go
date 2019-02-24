package main

import (
	"fmt"
	"strconv"
)

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

type Human struct {
	name string
	age int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func printStruct() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}


//反射推断接口中的类型
type Element interface {}
type List []  Element

type Student struct {
	name string
	age int
	school string
}

func (stu Student) String() string {
	return "name: " + stu.name + " school: " + stu.school + " age: " +  strconv.Itoa(stu.age)
}

func printInterfaceType()  {
	list :=make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Student{"heshuai", 20, "heYangSchool"}

	for index, element := range list {
		switch value :=  element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Student:
			fmt.Printf("list[%d] is a Student and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}

func main() {
	//printMap()
	//printSlice()
	//printStruct()
	printInterfaceType()
}
