package src

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	pNode := &ListNode{1, nil}
	pNode.nextNode = &ListNode{2, nil}
	pNode.nextNode.nextNode = &ListNode{3, nil}
	pNode.nextNode.nextNode.nextNode = &ListNode{4, nil}
	pNode.nextNode.nextNode.nextNode.nextNode = &ListNode{5, nil}
	DeleteNode(pNode, pNode)
}
