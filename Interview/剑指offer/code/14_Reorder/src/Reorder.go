package src

import (
	"fmt"
)

func ReorderOddEven(array []int) {
	if array == nil || len(array) == 0 {
		return
	}

	start := 0
	end := len(array) - 1

	for {
		if start > end {
			break
		}
		// 头指针向后移动，直到指向偶数
		for {
			if (array[start] & 0x1) == 0 {
				break
			}
			start++

		}

		// 尾指针向后移动，直到指向奇数
		for {
			if (array[end] & 0x1) != 0 {
				break
			}
			end--

		}
		if start < end {
			array[end], array[start] = array[start], array[end]
		}

	}

	fmt.Println(array)
}

func Reorder(array []int) {
	if array == nil || len(array) == 0 {
		return
	}

	start := 0
	end := len(array) - 1

	for {
		if start > end {
			break
		}
		for {
			if !IsEven(array[start]) {
				break
			}
			start++

		}
		for {
			if IsEven(array[end]) {
				break
			}
			end--
		}
		if start < end {
			array[end], array[start] = array[start], array[end]
		}
	}

	fmt.Println(array)
}

//分组标准：判断是不是偶数
func IsEven(n int) bool {
	return (n & 1) == 0
}
