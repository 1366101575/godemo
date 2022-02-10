package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {

	/*
		数组的特点
		( 1 ）数组创建完长度就固定了，不可以再追加元素。
		( 2 ）数组是值类型的，数组赋值或作为函数参数都是值拷贝。
		( 3 ）数组长度是数组类型的组成部分，[1O]int 和 [20]int 表示不同的类型。
		( 4 ）可以根据数组创建切片
	*/
	t.Run("test0", func(t *testing.T) {

		arr1 := [3]int{1, 2}
		arr2 := [...]int{5, 6}
		fmt.Printf("arr1:%v arr2:%v\n", arr1, arr2)

		//数组值传递
		changeArr := func(arr [3]int) {
			arr[2] = 999
			fmt.Printf("changeArr arr:%v\n", arr)
		}

		changeArr(arr1)

		for i, v := range arr1 {
			fmt.Printf("for range arr1 i:%d v:%d\n", i, v)
		}

		alen := len(arr2)
		for i := 0; i < alen; i++ {
			fmt.Printf("for arr2 i:%d v:%d\n", i, arr2[i])
		}

		fmt.Println("done")
	})
}
