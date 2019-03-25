package stackqueue

import (
	"fmt"
	"testing"

	"github.com/linthan/zuosuanfan/1chapter/minstatck"
)

func TestQueue(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2, 1}
	s := &StackQueue{
		stackPop:  &minstatck.Stack{},
		stackPush: &minstatck.Stack{},
	}
	for _, item := range numbers {
		s.Add(item)

	}
	for range numbers {
		fmt.Println("-------------")
		fmt.Println(s.Poll())
		fmt.Println("-------------")
	}
}
