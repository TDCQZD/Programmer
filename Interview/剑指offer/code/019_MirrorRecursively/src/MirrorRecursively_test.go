package src

import (
	"testing"
)

func TestMirrorRecursively(t *testing.T) {
	pRoot1 := NewBinaryTreeNode(8)
	pRoot2 := NewBinaryTreeNode(6)
	pRoot3 := NewBinaryTreeNode(10)
	pRoot4 := NewBinaryTreeNode(5)
	pRoot5 := NewBinaryTreeNode(7)
	pRoot6 := NewBinaryTreeNode(9)
	pRoot7 := NewBinaryTreeNode(11)

	pRoot1.Left = pRoot2
	pRoot1.Right = pRoot3

	pRoot2.Left = pRoot4
	pRoot2.Right = pRoot5

	pRoot3.Left = pRoot6
	pRoot3.Right = pRoot7
	PreOrder(pRoot1)
	MirrorRecursively(pRoot1)
	PreOrder(pRoot1)

}
