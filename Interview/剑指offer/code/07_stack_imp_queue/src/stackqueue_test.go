package src

import (
	"fmt"
	"testing"
)

func TestStackAndQueue(t *testing.T) {
	// testQueueWithTwoStacks(t)
	// testQueueOpera(t)
	testStackWithQueues(t)
}

func testQueueOpera(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	for {
		if queue.IsEmpty() {
			break
		}
		res := queue.Pop()
		fmt.Println(res)
	}
}
func testQueueWithTwoStacks(t *testing.T) {
	queue := NewMyQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	queue.Push(6)
	for {
		if queue.IsEmpty() {
			break
		}
		res := queue.Pop()
		fmt.Println(res)
	}
}

func testStackWithQueues(t *testing.T) {

	stack := NewMyStact()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)


	for {
		if stack.IsEmpty() {
			break
		}
		res := stack.Pop()
		fmt.Println(res)
	}
}
