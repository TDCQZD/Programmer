package src

type MyStack struct {
	queue1, queue2 *Queue
}

func NewMyStact() MyStack {
	return MyStack{queue1: NewQueue(), queue2: NewQueue()}
}

/* Push
入栈规则：
插入到非空队列，如果均为空则插入到queue1中
*/
func (ms *MyStack) Push(num int) {

	if !ms.queue1.IsEmpty() {
		ms.queue1.Push(num)
	}
	if !ms.queue2.IsEmpty() {
		ms.queue2.Push(num)
	}
	if ms.queue1.IsEmpty() && ms.queue2.IsEmpty() {
		ms.queue1.Push(num)
	}

}

/* Pop
两个队列满足：一队列为空，一队列非空。
出栈规则：
判断非空队列中元素的个数是否为1，如果等于1，则出队列;
否则将非空队列中的元素出队列并放入空队列中，直到非空队列中的元素只有一个，然后将剩余的一个元素出队列。
*/
func (ms *MyStack) Pop() int {

	if ms.queue2.IsEmpty() {
		if ms.queue1.Len() == 1 {
			return ms.queue1.Pop()
		} else {
			for {
				if ms.queue1.Len() == 1 {
					break
				}
				ms.queue2.Push(ms.queue1.Pop())
			}

			return ms.queue1.Pop()
		}
	}

	if ms.queue1.IsEmpty() {
		if ms.queue2.Len() == 1 {
			return ms.queue2.Pop()
		}
		for {
			if ms.queue2.Len() == 1 {
				break
			}
			ms.queue1.Push(ms.queue2.Pop())
		}
		return ms.queue2.Pop()
	}

	return 0

}

func (ms *MyStack) IsEmpty() bool {
	return ms.queue1.IsEmpty() && ms.queue2.IsEmpty()
}
