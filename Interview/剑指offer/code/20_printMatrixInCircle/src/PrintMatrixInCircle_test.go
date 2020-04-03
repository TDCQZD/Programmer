package src

import (
	"testing"
)

func TestPrintMatrixCircle(t *testing.T) {
	array := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	PrintMatrixInClockwisely(array, 4, 4)
}
