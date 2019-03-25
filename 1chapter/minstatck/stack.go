package minstatck

import "errors"

//Pop  弹出数据
func (s *Stack) Pop() (ret int, err error) {
	if len(s.data) == 0 {
		err = errors.New("has nothing to pop")
		return
	}
	ret = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return
}

//Push  压入数据
func (s *Stack) Push(ret int) {
	s.data = append(s.data, ret)
	return
}

//Peek 查看数据
func (s *Stack) Peek() (ret int, err error) {
	if len(s.data) == 0 {
		err = errors.New("has nothing to peek")
		return
	}
	ret = s.data[len(s.data)-1]
	return
}

//长度
func (s *Stack) Len() int {
	return len(s.data)
}
