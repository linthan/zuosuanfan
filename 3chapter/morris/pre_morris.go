package morris

import (
	"fmt"

	"github.com/linthan/zuosuanfan/3chapter/recurse"
)

//Node 节点
type Node = recurse.Node

//In 中序遍历
func In(head *Node) {
	if head == nil {
		return
	}
	cur1 := head
	var cur2 *Node
	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				cur2.Right = cur1
				cur1 = cur1.Left
				continue
			} else {
				//恢复遍历
				cur2.Right = nil
			}
		}
		fmt.Printf(cur1.Value + " ")
		cur1 = cur1.Right
	}
}

//Pre 先序遍历
func Pre(head *Node) {
	if head == nil {
		return
	}
	cur1 := head
	var cur2 *Node
	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				cur2.Right = cur1
				fmt.Printf(cur1.Value + " ")
				cur1 = cur1.Left
				continue
			} else {
				//恢复遍历
				cur2.Right = nil
			}
		} else {
			fmt.Printf(cur1.Value + " ")
		}
		cur1 = cur1.Right
	}
}

//Post 后续遍历
func Post(head *Node) {
	if head == nil {
		return
	}
	cur1 := head
	var cur2 *Node
	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				cur2.Right = cur1
				cur1 = cur1.Left
				continue
			} else {
				//恢复遍历
				cur2.Right = nil
				printEdge(cur1.Left)
			}
		}
		cur1 = cur1.Right
	}
	printEdge(head)
}

func printEdge(head *Node) {
	tail := reverseEdge(head)
	cur := tail
	for cur != nil {
		fmt.Printf(cur.Value + " ")
		cur = cur.Right
	}
	reverseEdge(tail)

}

//reverseEdge 逆序
func reverseEdge(from *Node) *Node {
	var pre, next *Node
	for from != nil {
		next = from.Right
		from.Right = pre
		pre = from
		from = next
	}
	return pre
}
