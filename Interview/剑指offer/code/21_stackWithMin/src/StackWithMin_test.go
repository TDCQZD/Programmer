package src

import (
	"fmt"
	"testing"
)

func TestStackWithMin(t *testing.T) {
	stack := NewStackWithMin()

	stack.push(3)
	stack.push(4)
	stack.push(2)
	stack.push(1)
	min := stack.min()
	fmt.Println("栈的最小元素：", min)
}
