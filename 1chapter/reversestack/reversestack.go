package reversestack

import (
	"github.com/linthan/zuosuanfan/1chapter/minstatck"
)

//ReverseWrapper Wrapper
func ReverseWrapper(s *minstatck.Stack) (b *minstatck.Stack) {
	b = &minstatck.Stack{}
	for {
		a := Reverse(s)
		b.Push(a)
		if s.Len() == 0 {
			return
		}
	}
}

//Reverse  逆序栈
func Reverse(s *minstatck.Stack) int {
	a, err := s.Pop()
	if s.Len() == 0 {
		return a
	}
	b := Reverse(s)
	if err == nil {
		s.Push(a)
	}
	return b

}
