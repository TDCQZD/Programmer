package src

import "fmt"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func NewBinaryTreeNode(val int) *BinaryTree {
	pRoot := &BinaryTree{}
	pRoot.Val = val
	pRoot.Left = nil
	pRoot.Right = nil

	return pRoot
}
func MirrorRecursively(pNode *BinaryTree) {
	if pNode == nil || pNode == nil || pNode == nil {
		return
	}

	tmp := pNode.Left
	pNode.Left = pNode.Right
	pNode.Right = tmp

	if pNode.Left != nil {
		MirrorRecursively(pNode.Left)
	}

	if pNode.Right != nil {
		MirrorRecursively(pNode.Right)
	}
}

//前序遍历
func PreOrder(tree *BinaryTree) {

	if tree != nil {
		fmt.Println(tree.Val)
		PreOrder(tree.Left)
		PreOrder(tree.Right)
	}
}
