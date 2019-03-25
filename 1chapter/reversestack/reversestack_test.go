package reversestack

import (
	"fmt"
	"testing"

	"github.com/linthan/zuosuanfan/1chapter/minstatck"
)

func TestS(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2, 1}
	s := &minstatck.Stack{}
	for _, item := range numbers {
		s.Push(item)

	}
	//逆序栈
	s = ReverseWrapper(s)
	for range numbers {
		fmt.Println(s.Pop())
	}
}
