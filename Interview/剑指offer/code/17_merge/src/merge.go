package src

type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge(pHead1, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	} else if pHead2 == nil {
		return pHead1
	}

	pMergeHead := &ListNode{}

	if pHead1.Val < pHead2.Val {
		pMergeHead = pHead1
		pMergeHead.Next = Merge(pHead1.Next, pHead2)
	} else {
		pMergeHead = pHead2
		pMergeHead.Next = Merge(pHead1, pHead2.Next)
	}

	return pMergeHead
}
