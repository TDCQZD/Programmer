package src

func ReverseList(pHead *ListNode) *ListNode {
	pPeverseHead := &ListNode{} ///反转链表头结点
	pNode := pHead       // 当前遍历的节点
	pPrev := &ListNode{} // 当前遍历的前一个节点

	for pNode != nil {
		pNext := pNode.Next // 当前遍历的后一个节点
		if pNext == nil {//如果当前结点的下一个结点为空，那么反转链表的头结点就是当前结点。
			pPeverseHead = pNode
		}

		pNode.Next = pPrev
		pPrev = pNode
		pNode = pNext
	}
	return pPeverseHead
}
