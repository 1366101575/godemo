package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 3, 5, 7, 8}

	count := 0
	for _, val := range arr {
		count += val
	}
	fmt.Printf("\n%v count:%d\n", arr, count)

	var find [][2]int
	arrLen := len(arr)
	for i1 := 0; i1 < arrLen-1; i1++ {
		for i2 := i1 + 1; i2 < arrLen; i2++ {
			if arr[i1]+arr[i2] == 8 {
				find = append(find, [2]int{i1, i2})
			}
		}
	}
	fmt.Printf("\n%v find equal 8 index:%v\n", arr, find)
}
