package src

import (
	"fmt"
	"testing"
)

func TestListResver(t *testing.T) {

	// testListReversing(t)
	testListReversing_Recursively(t)
	// testListReversing_Iteratively(t)
	// testReverseList(t)
}

func testListReversing(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}

	tmp := ListReversingly(head)
	recursive(tmp) // 递归遍历链表
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}

	array := []int{1, 2, 3, 4, 5}
	recursiveArray(array)
}
func testListReversing_Iteratively(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	ListReversingly_Iteratively(head)

}
func testListReversing_Recursively(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	ListReversingly_Recursively(head)

}

func testReverseList(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	// recursive(head)
	recursive(ReverseList(head))

}
