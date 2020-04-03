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

func HasSubTree(pRoot1, pRoot2 *BinaryTreeNode) (flag bool) {
	flag = false
	if pRoot1 != nil && pRoot2 != nil {
		if pRoot1.Val == pRoot2.Val {
		
			flag = DoesTree1HaveTree2(pRoot1, pRoot2)
		}
		if !flag {
			
			flag = HasSubTree(pRoot1.Left, pRoot2)
		}
		if !flag {
			
			flag = HasSubTree(pRoot1.Right, pRoot2)
		}
	}
	return flag

}

func DoesTree1HaveTree2(pRoot1, pRoot2 *BinaryTreeNode) bool {

	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}

	if pRoot1.Val != pRoot2.Val {
		return false
	}

	return DoesTree1HaveTree2(pRoot1.Left, pRoot2.Left) && DoesTree1HaveTree2(pRoot1.Right, pRoot2.Right)
}
