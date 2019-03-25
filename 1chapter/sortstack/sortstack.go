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

//SortWrapper1 Wrapper
func SortWrapper1(s *minstatck.Stack) {
	help := &minstatck.Stack{}
	for s.Len() != 0 {
		ret, _ := s.Pop()
		for help.Len() != 0 {
			p, _ := help.Peek()
			//一定要带上等于 要不然会在上层发生死循环
			if p <= ret {
				break
			}
			tmp, _ := help.Pop()
			s.Push(tmp)
		}
		help.Push(ret)
	}

	for help.Len() != 0 {
		tmp, _ := help.Pop()
		s.Push(tmp)
	}

}
