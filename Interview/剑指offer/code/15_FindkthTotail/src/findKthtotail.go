package src

type ListNode struct {
	val  int
	Next *ListNode
}

func FindKthToTail(pHeadNode *ListNode, k int) *ListNode {
	if pHeadNode == nil || k == 0 {
		return nil
	}
	pHead := pHeadNode
	pBehind := &ListNode{}

	for i := 0; i < k-1; i++ {
		if pHead.Next != nil {
			pHead = pHead.Next
		} else {
			return nil
		}

	}
	pBehind = pHeadNode

	for pHead.Next != nil {
		pHead = pHead.Next
		pBehind = pBehind.Next
	}
	return pBehind
}
