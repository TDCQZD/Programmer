package src

// Queue 利用 栈 实现的队列
type MyQueue struct {
	stack1, stack2 *Stack
}

// NewQueue 构造函数初始化数据结构。
func NewMyQueue() MyQueue {
	return MyQueue{stack1: NewStack(), stack2: NewStack()}
}

// 入队列
func (queue *MyQueue) Push(num int) {
	queue.stack1.Push(num)
}

// 出队列 重点!!!!!!
func (queue *MyQueue) Pop() int {
	if queue.stack2.IsEmpty() {
		for !queue.stack1.IsEmpty() {
			queue.stack2.Push(queue.stack1.Pop())
		}
	}

	return queue.stack2.Pop()
}

// 判断队列是否为空
func (queue *MyQueue) IsEmpty() bool {
	return queue.stack1.IsEmpty() && queue.stack2.IsEmpty()
}
