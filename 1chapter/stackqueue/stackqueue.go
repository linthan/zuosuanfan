package stackqueue

import (
	"github.com/linthan/zuosuanfan/1chapter/minstatck"
)

//StackQueue
type StackQueue struct {
	stackPush *minstatck.Stack
	stackPop  *minstatck.Stack
}

//Add 添加
func (s *StackQueue) Add(ret int) {
	if s.stackPop.Len() != 0 {
		for {
			a, err := s.stackPop.Pop()
			if err != nil {
				break
			}
			s.stackPush.Push(a)
		}
	}
	s.stackPush.Push(ret)
	return
}

//Poll
func (s *StackQueue) Poll() (ret int, err error) {
	if s.stackPop.Len() == 0 {
		for {
			a, err := s.stackPush.Pop()
			if err != nil {
				break
			}
			s.stackPop.Push(a)
		}
	}
	ret, err = s.stackPop.Pop()
	return
}
