package minstatck

//Stack 栈顶元素
type Stack struct {
	data []int
}

//MinStack 需要的minStack
type MinStack struct {
	normal *Stack //正常的stack
	min    *Stack //最小stack
}

//Pop 弹出元素
func (m *MinStack) Pop() (ret int, err error) {
	ret, err = m.normal.Pop()
	if err != nil {
		return
	}
	if pdata, _ := m.min.Peek(); pdata == ret {
		m.min.Pop()
	}
	return

}

//Push 压入元素
func (m *MinStack) Push(ret int) {
	m.normal.Push(ret)
	if pdata, _ := m.min.Peek(); m.min.Len() == 0 || pdata >= ret {
		m.min.Push(ret)
	}
	return
}

//Push 压入元素
func (m *MinStack) GetMin() (ret int, err error) {
	ret, err = m.min.Peek()
	if err != nil {
		return
	}
	return
}
