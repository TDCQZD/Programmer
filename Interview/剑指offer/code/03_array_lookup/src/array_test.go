package src

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	testArray(t)
}

/*测试用例：

二维数组中包含査找的数字（査找的数字是数组中的最大值和最小
值，査找的数字介丁•数组中的最大值和最小值之间）。

二维数组中没有査找的数字（査找的数字大于数组中的最大值，査
找的数字小于数组中的最小值，查找的数字在数组的最大值和最小
值之间但数组中没有这个数字

特殊输入测试（输入数组）。
*/
func testArray(t *testing.T) {
	matrix := [][]int{
		[]int{1, 2, 8, 9},
		[]int{2, 4, 9, 12},
		[]int{4, 7, 10, 13},
		[]int{6, 8, 11, 15},
	}
	// fmt.Println("len:",len(matrix))
	// flag, row, cloumn := FindMatrix(matrix, 7)
	flag, row, cloumn := FindMatrix2(matrix, 4, 4, 16)
	if flag {
		fmt.Printf("array[%d][%d] = %d \n", row, cloumn, matrix[row][cloumn])
	} else {
		fmt.Println("该数字不存在！")
	}

}
