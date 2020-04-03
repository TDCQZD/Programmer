package src

type StackWithMin struct {
	m_data *Stack
	m_min  *Stack
}

func NewStackWithMin() *StackWithMin {
	return &StackWithMin{m_data: NewStack(), m_min: NewStack()}
}
func (swm *StackWithMin) push(num int) {
	swm.m_data.Push(num)
	if swm.m_min.Len() == 0 || num < swm.m_min.Top() {
		swm.m_min.Push(num)
	} else {
		swm.m_min.Push(swm.m_min.Top())
	}

}
func (swm *StackWithMin) pop() {
	if swm.m_data.Len() > 0 && swm.m_min.Len() > 0 {
		swm.m_data.Pop()
		swm.m_min.Pop()
	}
}

func (swm *StackWithMin) top() int {
	return swm.m_data.Top()
}
func (swm *StackWithMin) min() int {
	if swm.m_data.Len() > 0 && swm.m_min.Len() > 0 {
		return swm.m_min.Top()
	}
	return 0
}

func (swm *StackWithMin) empty() bool {
	return swm.m_data.IsEmpty()
}
func (swm *StackWithMin) size() int {
	return swm.m_data.Len()
}
