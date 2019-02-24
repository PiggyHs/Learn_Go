package main

import "fmt"

//冒泡排序
func bubbleSort(arr []int, size int) []int{

	for i := 0; i < size -1; i++ {
		for j := 0; j < size - 1 - i; j++ {
			if arr[j] > arr[j + 1] {
				var temp  = arr[j]
				arr[j] = arr[j + 1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func main()  {
	arr := []int {9,8,7,6,5,4,3,2,1,10}

	arrayList := bubbleSort(arr, len(arr))

	for i := 0; i<len(arrayList); i++  {
		fmt.Println(arrayList[i])
	}
}