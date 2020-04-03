package src

func VerifySquenceOfBST(queue []int, length int) bool {
	if queue == nil || len(queue) <= 0 {
		return false
	}
	root := queue[length-1]

	// 在二叉搜素树中左子树的节点小于根节点
	i := 0
	for ; i < length-1; i++ {
		if queue[i] > root {
			break
		}
	}
	// 在二叉搜素树中右子树的节点大于根节点
	j := i
	for ; j < length-1; j++ {
		if queue[j] < root {
			return false
		}
	}

	// 判断左子树是不是二叉搜素树
	left := true
	if i > 0 {
		left = VerifySquenceOfBST(queue[:i], i)
	}
	// 判断右子树是不是二叉搜素树
	right := true
	if i < length-1 {
		right = VerifySquenceOfBST(queue[i:length-1], length-i-1)
	}
	return left && right
}
