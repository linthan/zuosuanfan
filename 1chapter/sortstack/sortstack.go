package sortstack

import (
	"github.com/linthan/zuosuanfan/1chapter/minstatck"
)

//SortWrapper Wrapper
func SortWrapper(s *minstatck.Stack) (b *minstatck.Stack) {
	b = &minstatck.Stack{}
	for {
		a := removeMaxValue(s)
		b.Push(a)
		if s.Len() == 0 {
			return
		}
	}
}

//removeMaxValue  删除最大值
func removeMaxValue(s *minstatck.Stack) int {
	a, _ := s.Pop()
	if s.Len() == 0 {
		return a
	}
	b := removeMaxValue(s)
	if a < b {
		s.Push(a)
		return b
	} else {
		s.Push(b)
		return a
	}
}
