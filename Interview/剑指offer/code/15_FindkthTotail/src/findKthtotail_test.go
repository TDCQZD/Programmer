package src

import (
	"fmt"
	"testing"
)

func TestFindKthTotail(t *testing.T) {
	pListNode := &ListNode{val: 1, Next: nil}
	pListNode.Next = &ListNode{val: 2, Next: nil}
	pListNode.Next.Next = &ListNode{val: 3, Next: nil}
	pListNode.Next.Next.Next = &ListNode{val: 4, Next: nil}
	pListNode.Next.Next.Next.Next = &ListNode{val: 5, Next: nil}

	k:=3

	kthNode := FindKthToTail(pListNode, k)
	fmt.Printf("链表中倒数第 %d个结点是：%v \n", k, kthNode)
}
