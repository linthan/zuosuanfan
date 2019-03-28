package recurse

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/stacks/arraystack"
)

const (
	cur = "12!3!#!#!#!"
)

//Node 节点
type Node struct {
	Value string `json:"value"`
	Left  *Node  `json:"left"`
	Right *Node  `json:"right"`
}

//PreOrderRecur 先序遍历
func PreOrderRecur(head *Node) {
	if head == nil {
		return
	}
	fmt.Println(head.Value)
	PreOrderRecur(head.Left)
	PreOrderRecur(head.Right)
}

//CreateTree 创建树
func CreateTree(preStr string) *Node {
	values := strings.Split(preStr, "!")
	stack := arraystack.New()
	for i := len(values) - 1; i >= 0; i-- {
		stack.Push(values[i])
	}
	return reconPreOrder(stack)
}

//reconPreOrder 重新创建
func reconPreOrder(stack *arraystack.Stack) *Node {
	value, ok := stack.Pop()
	if !ok || value == nil || value.(string) == "#" {
		return nil
	}
	head := &Node{
		Value: value.(string),
	}
	head.Left = reconPreOrder(stack)
	head.Right = reconPreOrder(stack)
	return head
}

func serialByPre(head *Node) string {
	if head == nil {
		return "#!"
	}
	res := head.Value + "!"
	res += serialByPre(head.Left)
	res += serialByPre(head.Right)
	return res
}
