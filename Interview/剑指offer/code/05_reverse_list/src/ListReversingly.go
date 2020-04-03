package src

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归反转链表
func ListReversingly_Recursively(head *ListNode) {

	if head == nil {
		return
	}
	if head.Next != nil {
		ListReversingly_Recursively(head.Next)
	}

	fmt.Println(head.Val)
}

// 使用两个常量辅助反转链表——最佳实践
func ListReversingly(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	dummy := head
	tmp := head

	for head != nil && head.Next != nil {
		dummy = head.Next
		head.Next = dummy.Next
		dummy.Next = tmp
		tmp = dummy
	}

	return dummy

}

// 递归遍历链表
func recursive(head *ListNode) {
	if head == nil {
		return
	}

	fmt.Println(head.Val)
	recursive(head.Next)
}

// 递归遍历数组
func recursiveArray(array []int) {
	if len(array) == 0 {
		return
	}

	fmt.Println(array[0])
	recursiveArray(array[1:])
}
