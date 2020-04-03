package src

import (
	"fmt"
)

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}
type QueueTree struct {
	nodes []*BinaryTreeNode
}

// NewQueue 返回
func NewQueue() *QueueTree {
	return &QueueTree{nodes: []*BinaryTreeNode{}}
}

// Push 把 n 放入 队列
func (s *QueueTree) Push(n *BinaryTreeNode) {
	s.nodes = append(s.nodes, n)
}

// Pop 从 s 中取出最后放入 队列 的值
func (s *QueueTree) Pop() *BinaryTreeNode {
	res := s.nodes[0]
	s.nodes = s.nodes[1:]
	return res
}

func (s *QueueTree) Top() *BinaryTreeNode {
	res := s.nodes[0]
	return res
}

// Len 返回 s 的长度
func (s *QueueTree) Len() int {
	return len(s.nodes)
}

// IsEmpty 反馈 s 是否为空
func (s *QueueTree) IsEmpty() bool {
	return s.Len() == 0
}

func PrintToptoBottom(root *BinaryTreeNode) {

	if root == nil {
		return
	}

	queue := NewQueue()
	queue.Push(root)
	for {
		if queue.IsEmpty()  {
			break
		}
		pNode := queue.Top()
		queue.Pop()
		fmt.Println(pNode.Val)

		if pNode.Left != nil {
			queue.Push(pNode.Left)
		}
		if pNode.Right != nil {
			queue.Push(pNode.Right)
		}

	}
}
