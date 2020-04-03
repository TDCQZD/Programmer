package src

import (
	"fmt"
)

func PrintMatrixInClockwisely(matrix [][]int, cloumns, rows int) {
	if matrix == nil || cloumns <= 0 || rows <= 0 {
		return
	}
	start := 0
	for cloumns > start*2 && rows > start*2 {
		PrintMatrixCircle(matrix, cloumns, rows, start)
		start++
	}
}
func PrintMatrixCircle(matrix [][]int, cloumns, rows, start int) {
	endX := cloumns - 1 - start
	endY := rows - 1 - start
	// 从左到右打印一行
	for i := start; i <= endX; i++ {
		num := matrix[start][i]
		fmt.Println(num)
	}

	// 从上到下打印一列
	if start < endY {
		for i := start + 1; i <= endY; i++ {
			num := matrix[i][endX]
			fmt.Println(num)
		}
	}

	// 从右到左打印一行
	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			num := matrix[endY][i]
			fmt.Println(num)
		}
	}

	// 从下到上打印一列
	if start < endX && start < endY-1 {
		for i := endY - 1; i >= start+1; i-- {
		
			num := matrix[i][start]
			fmt.Println(num)
		}
	}

}
