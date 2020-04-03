package src

import "fmt"

type ListNode struct {
	val      int
	nextNode *ListNode
}

//删除链表指定节点
func DeleteNode(pHeadNode, pToBeDeleted *ListNode) {
	if pHeadNode == nil || pToBeDeleted == nil {
		return
	}
	// 待删除的结点不是尾结点
	if pToBeDeleted.nextNode != nil {

		pNode := pToBeDeleted.nextNode
		pToBeDeleted.val = pNode.val
		pToBeDeleted.nextNode = pNode.nextNode
	} else if pHeadNode == pToBeDeleted { //链表只有一个节点，删除头结点(也是尾结点)
		pHeadNode = nil
		fmt.Println("空链表！")
	} else { //链表有多个节点，删除尾结点
		pNode := pHeadNode
		for pNode.nextNode != pToBeDeleted {
			pNode = pNode.nextNode
		}
		pNode.nextNode = nil
	}
	recursive(pHeadNode)
}

// 递归遍历链表
func recursive(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.val)
	recursive(head.nextNode)
}
