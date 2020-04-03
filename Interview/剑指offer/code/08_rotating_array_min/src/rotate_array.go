package src

import (
	"fmt"
)

func Min(array []int) {

	if array == nil || len(array) <= 0 {
		fmt.Println("数组为空！")
		return
	}

	index1 := 0
	index2 := len(array) - 1

	indexMiddle := index1

	for array[index1] >= array[index2] {
		if index2-index1 == 1 { //当两个元素相邻时，循环结束
			indexMiddle = index2
			break
		}
		indexMiddle = (index1 + index2) / 2

		//如果下标index1、index2和indexMiddle指向的值相等，则顺序查找
		if array[index1] == array[index2] && array[index1] == array[indexMiddle] {
			MinInorder(array)
			return
		}

		if array[index1] <= array[indexMiddle] {
			index1 = indexMiddle
		} else if array[index2] >= array[indexMiddle] {
			index2 = indexMiddle
		}
	}
	fmt.Println("旋转数组的最小数字:", array[indexMiddle])
}

func MinInorder(array []int) {
	fmt.Println("旋转数组顺序查找")
	middle := array[0]
	for _, value := range array {
		if middle <= value {
			middle = value
		}

	}
	fmt.Println("旋转数组的最小数字:", middle)
}
