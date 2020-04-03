package src

import (
	"fmt"
	"testing"
)

func TestVerifySquenceOfBST(t *testing.T) {
	// queue := []int{5, 7, 6, 9, 11, 10, 8}
	// res := VerifySquenceOfBST(queue,7)
	queue := []int{7,4,6,5}
	res := VerifySquenceOfBST(queue,4)
	fmt.Println(res)
}
