package reversestack

import (
	"github.com/linthan/zuosuanfan/1chapter/minstatck"
)

//ReverseWrapper Wrapper
func ReverseWrapper(s *minstatck.Stack) {
	a := Reverse(s)
	s.Push(a)
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
