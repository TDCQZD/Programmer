package src

import (
	"testing"
)

func Test(t *testing.T) {
	test2(t)
}

// 正常情况
func test(t *testing.T) {
	//原数组：[1,2,3,4,5]
	//旋转数组
	array := []int{3, 4, 5, 1, 2}
	Min(array)

}

//特殊情况：空数组
func test1(t *testing.T) {
	Min([]int{})
}

//特殊情况：第一个数字、中间数字和最后一个数字相等
func test2(t *testing.T) {
	//原数组：[0,1,1,1,1]
	//旋转数组
	// array := []int{1, 0, 1, 1, 1}
	array := []int{1, 1, 1, 0, 1}

	Min(array)

}
