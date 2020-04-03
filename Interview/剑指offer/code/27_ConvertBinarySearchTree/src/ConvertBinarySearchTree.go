package src

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewBinaryTreeNode(val int) *BinaryTreeNode {
	pRoot := &BinaryTreeNode{}
	pRoot.Val = val
	pRoot.Left = nil
	pRoot.Right = nil

	return pRoot
}
func ConvertBinarySearchTree(pRoot *BinaryTreeNode) *BinaryTreeNode {
	pLastNodeInlist := &BinaryTreeNode{} // 指向双向链表的尾结点
	ConvertNode(pRoot, &pLastNodeInlist)
	// 返回头结点
	pHeadList := pLastNodeInlist
	for {
		if pHeadList == nil || pHeadList.Left == nil {
			break
		}
		pHeadList = pHeadList.Left
	}
	return pHeadList
}

func ConvertNode(pNode *BinaryTreeNode, pLastNodeInlist **BinaryTreeNode) {
	if pNode == nil {
		return
	}
	pCurrentNode := pNode
	if pCurrentNode.Left != nil {
		ConvertNode(pCurrentNode.Left, pLastNodeInlist)
	}
	pCurrentNode.Left = *pLastNodeInlist
	if *pLastNodeInlist != nil {
		(*pLastNodeInlist).Right = pCurrentNode
	}
	*pLastNodeInlist = pCurrentNode
	
	if pCurrentNode.Right != nil {
		ConvertNode(pCurrentNode.Right, pLastNodeInlist)
	}
}
