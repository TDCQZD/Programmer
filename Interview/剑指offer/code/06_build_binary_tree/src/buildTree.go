package src

import (
	"fmt"
)

/*
前序遍历序列{ 1，2，4，7，3，5，6，8 }
中序遍历序列{ 4，7，2，1，5，3，8，6}，

重建出图中所示的二叉树并输出它的头结点。
*/
//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}

func buildTreeByPre(pre, in []int) *TreeNode {

	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return root
	}
	// 确定根节点的值位置
	idx := indexOf(root.Val, in)
	// 构建左子树
	root.Left = buildTreeByPre(pre[1:idx+1], in[:idx])
	// 构建右子树
	root.Right = buildTreeByPre(pre[idx+1:], in[idx+1:])

	return root

}
func buildTreeByPost(in, post []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: post[len(post)-1],
	}

	if len(in) == 1 {
		return root
	}
	// 确定根节点的值位置
	idx := indexOf(root.Val, in)
	fmt.Println("idx:", idx)
	// 构建左子树
	root.Left = buildTreeByPost(in[:idx], post[:idx])
	// 构建右子树
	root.Right = buildTreeByPost(in[idx+1:], post[idx:len(post)-1])

	return root

}

//前序遍历
func PreOrder(tree *TreeNode) {

	if tree != nil {
		fmt.Println(tree.Val)
		PreOrder(tree.Left)
		PreOrder(tree.Right)
	}
}

//中序遍历
func InfixOrder(tree *TreeNode) {

	if tree != nil {
		InfixOrder(tree.Left)
		fmt.Println(tree.Val)
		InfixOrder(tree.Right)
	}
}

//后序遍历
func PostOrder(tree *TreeNode) {

	if tree != nil {
		PostOrder(tree.Left)
		PostOrder(tree.Right)
		fmt.Println(tree.Val)
	}
}
