package minstatck

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2, 1}
	s := &MinStack{
		normal: &Stack{},
		min:    &Stack{},
	}
	for _, item := range numbers {
		s.Push(item)
		fmt.Println(s.GetMin())

	}
	fmt.Println("--------------------------------------------------------")
	for range numbers {
		fmt.Println("-------------")
		fmt.Println(s.GetMin())
		fmt.Println(s.Pop())
		fmt.Println("-------------")
	}
}
