package src

import (
	"fmt"
	"testing"
)

func TestIsPopOrder(t *testing.T) {
	nLength := 5
	push := []int{1, 2, 3, 4, 5}
	pop := []int{4, 5, 3, 2, 1}
	res := IsPopOrder(push, pop, nLength)
	fmt.Println(res)
}
