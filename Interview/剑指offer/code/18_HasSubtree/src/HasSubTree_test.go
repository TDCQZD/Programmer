package src

import (
	"fmt"
	"testing"
)

func TestHasSubTree(t *testing.T) {
	pRootA1 := NewBinaryTreeNode(8)
	pRootA2 := NewBinaryTreeNode(8)
	pRootA3 := NewBinaryTreeNode(7)
	pRootA4 := NewBinaryTreeNode(9)
	pRootA5 := NewBinaryTreeNode(2)
	pRootA6 := NewBinaryTreeNode(4)
	pRootA7 := NewBinaryTreeNode(7)

	pRootA1.Left = pRootA2
	pRootA1.Right = pRootA3

	pRootA2.Left = pRootA4
	pRootA2.Right = pRootA5

	pRootA5.Left = pRootA6
	pRootA5.Right = pRootA7

	pRootB1 := NewBinaryTreeNode(8)
	pRootB2 := NewBinaryTreeNode(9)
	pRootB3 := NewBinaryTreeNode(1)

	pRootB1.Left = pRootB2
	pRootB1.Right = pRootB3

	flag := HasSubTree(pRootA1, pRootB1)
	fmt.Println(flag)

}
