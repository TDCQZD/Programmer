package src

/*FindMatrix 二维数组查找*/
func FindMatrix(matrix [][]int, number int) (flag bool, row, cloum int) {
	if matrix != nil && len(matrix) != 0 {
		// 二维数组右上角数字
		i := 0                  // 行
		j := len(matrix[0]) - 1 //列

		for i < len(matrix) && j >= 0 {
			if matrix[i][j] == number {
				return true, i, j
			} else if matrix[i][j] > number {
				j--
			} else {
				i++
			}
		}
		return false, 0, 0
	} else {
		return false, 0, 0
	}

}

func FindMatrix2(matrix [][]int, rows int, columns int, number int) (flag bool, row, cloum int) {
	if matrix != nil && len(matrix) != 0 {
		row := 0
		cloum := columns - 1
		for row < rows && cloum >= 0 {
			if matrix[row][cloum] == number {
				return true, row, cloum
			} else if matrix[row][cloum] > number {
				cloum--
			} else {
				row++
			}
		}
		return false, 0, 0
	} else {
		return false, 0, 0
	}

}
