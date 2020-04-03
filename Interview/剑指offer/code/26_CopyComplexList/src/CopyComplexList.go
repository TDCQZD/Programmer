package src

import "fmt"

type ComplexListNode struct {
	m_Value    int
	m_pNext    *ComplexListNode
	m_pSibling *ComplexListNode
}

func CreateNode(nValue int) *ComplexListNode {

	pNode := &ComplexListNode{}
	pNode.m_Value = nValue
	pNode.m_pNext = nil
	pNode.m_pSibling = nil

	return pNode
}

func BuildNodes(pNode, pNext, pSibling *ComplexListNode) {
	if pNode != nil {
		pNode.m_pNext = pNext
		pNode.m_pSibling = pSibling
	}
}
func CopyComplexList(pHead *ComplexListNode) *ComplexListNode {
	CloneNodes(pHead)
	ConnectSiblingNodes(pHead)
	return ReconnectNode(pHead)
}

func PrintList(pHead *ComplexListNode) {
	pNode := pHead
	for {
		if pNode == nil {
			break
		}
		fmt.Printf("The value of this node is: %d.\n", pNode.m_Value)

		if pNode.m_pSibling != nil {
			fmt.Printf("The value of its sibling is: %d.\n", pNode.m_pSibling.m_Value)
		} else {
			fmt.Printf("This node does not have a sibling.\n")
		}
		fmt.Printf("\n")

		pNode = pNode.m_pNext
	}
}

// 复制节点
func CloneNodes(pHead *ComplexListNode) {
	pNode := pHead
	for {
		if pNode == nil {
			break
		}
		pCloned := &ComplexListNode{}
		pCloned.m_Value = pNode.m_Value
		pCloned.m_pNext = pNode.m_pNext
		pCloned.m_pSibling = nil

		pNode.m_pNext = pCloned
		pNode = pCloned.m_pNext
	}
}

// 复制节点的m_pSibling
func ConnectSiblingNodes(pHead *ComplexListNode) {
	pNode := pHead
	for {
		if pNode == nil {
			break
		}
		pCloned := pNode.m_pNext
		if pNode.m_pSibling != nil {
			pCloned.m_pSibling = pNode.m_pSibling.m_pNext
		}
		pNode = pCloned.m_pNext
	}
}

// 将长链表分为两个链表，奇数为的节点为原链表，偶数位为复制链表
func ReconnectNode(pHead *ComplexListNode) *ComplexListNode {
	pNode := pHead
	pClonedHead := &ComplexListNode{}
	pClonedNode := &ComplexListNode{}

	if pNode != nil {
		pClonedHead = pNode.m_pNext
		pClonedNode = pNode.m_pNext
		pNode.m_pNext = pClonedNode.m_pNext
		pNode = pNode.m_pNext
	}

	for {
		if pNode == nil {
			break
		}
		pClonedNode.m_pNext = pNode.m_pNext
		pClonedNode = pClonedNode.m_pNext
		pNode.m_pNext = pClonedNode.m_pNext
		pNode = pNode.m_pNext
	}
	return pClonedHead
}
