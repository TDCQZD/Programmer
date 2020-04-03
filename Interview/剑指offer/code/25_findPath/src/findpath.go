package src

import (
	"fmt"
)

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func FindPath(pRoot *BinaryTreeNode, expectedNum int) {
	if pRoot == nil {
		return
	}
	path := []int{}
	currentSum := 0
	SearchPath(pRoot, expectedNum, path, currentSum)
}

func SearchPath(pRoot *BinaryTreeNode, expectedNum int, path []int, currentSum int) {
	currentSum += pRoot.Val
	path = append(path, pRoot.Val)
	fmt.Println(path)

	// 如果是叶节点，并且路径上结点的和等于输入的值，打印出这条路径
	isLeaf := pRoot.Left == nil && pRoot.Right == nil

	if currentSum == expectedNum && isLeaf {	
		for _, value := range path {
			fmt.Println(value)
		}
	}

	// 如果不是叶节点，则遍历它的子节点
	if pRoot.Left != nil {
		SearchPath(pRoot.Left, expectedNum, path, currentSum)
	}
	if pRoot.Right != nil {
		SearchPath(pRoot.Right, expectedNum, path, currentSum)
	}

	// 在返回到父节点之前，在路径上删除当前节点，并在currentSum中减去当前结点的值
	currentSum -= pRoot.Val
	path = path[:len(path)-1]
}
