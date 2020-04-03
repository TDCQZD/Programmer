package src

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	head1 := &ListNode{1, nil}
	head1.Next = &ListNode{3, nil}
	head1.Next.Next = &ListNode{5, nil}

	head2 := &ListNode{2, nil}
	head2.Next = &ListNode{4, nil}
	head2.Next.Next = &ListNode{6, nil}

	head := Merge(head1, head2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
