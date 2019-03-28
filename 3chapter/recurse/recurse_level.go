package recurse

import (
	"strings"

	"github.com/emirpasic/gods/lists/singlylinkedlist"
)

//serialByLevel 通过层序来遍历二叉树
func serialByLevel(head *Node) string {
	if head == nil {
		return "#!"
	}
	res := head.Value + "!"
	queue := singlylinkedlist.New()
	queue.Append(head)
	for !queue.Empty() {
		value, _ := queue.Get(0)
		head = value.(*Node)
		queue.Remove(0)
		if head.Left != nil {
			res += head.Left.Value + "!"
			queue.Add(head.Left)
		} else {
			res += "#!"
		}
		if head.Right != nil {
			res += head.Right.Value + "!"
			queue.Add(head.Right)
		} else {
			res += "#!"
		}
	}
	return res
}

func reconByLevelString(levelStr string) *Node {
	values := strings.Split(levelStr, "!")
	index := 0
	head := generateNodeByString(values[index])
	index++
	queue := singlylinkedlist.New()
	queue.Append(head)
	var tmp *Node
	for !queue.Empty() {
		value, _ := queue.Get(0)
		tmp = value.(*Node)
		queue.Remove(0)
		tmp.Left = generateNodeByString(values[index])
		index++
		tmp.Right = generateNodeByString(values[index])
		index++
		if tmp.Left != nil {
			queue.Append(tmp.Left)
		}
		if tmp.Right != nil {
			queue.Append(tmp.Right)
		}
	}
	return head
}

func generateNodeByString(st string) *Node {
	if st == "#" {
		return nil
	}
	return &Node{
		Value: st,
	}
}
