package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 5, 6, 1, 7, 9, 8, 10}
	sort.Ints(a)

	for i, v := range a {
		fmt.Println(i, v)
	}
}