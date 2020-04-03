package src

import "testing"

func TestFindPath(t *testing.T) {
	rootNode := &BinaryTreeNode{Val: 10}
	pNode1 := &BinaryTreeNode{Val: 5}
	pNode2 := &BinaryTreeNode{Val: 12}
	pNode3 := &BinaryTreeNode{Val: 4}
	pNode4 := &BinaryTreeNode{Val: 7}

	rootNode.Left = pNode1
	rootNode.Right = pNode2

	pNode1.Left = pNode3
	pNode1.Right = pNode4

	FindPath(rootNode, 22)
}
