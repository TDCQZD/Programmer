package src

import (
	"testing"
)

func TestPrintToptoBottom(t *testing.T) {
	rootNode := &BinaryTreeNode{Val: 8}
	pNode1 := &BinaryTreeNode{Val: 6}
	pNode2 := &BinaryTreeNode{Val: 10}
	pNode3 := &BinaryTreeNode{Val: 5}
	pNode4 := &BinaryTreeNode{Val: 7}
	pNode5 := &BinaryTreeNode{Val: 9}
	pNode6 := &BinaryTreeNode{Val: 11}

	rootNode.Left = pNode1
	rootNode.Right = pNode2

	pNode1.Left = pNode3
	pNode1.Right = pNode4

	pNode2.Left = pNode5
	pNode2.Right = pNode6

	PrintToptoBottom(rootNode)
}
