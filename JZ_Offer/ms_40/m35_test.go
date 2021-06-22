package ms_40

import (
	"testing"
)

// 剑指 Offer 35. 复杂链表的复制

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 深度优先搜索 -> 递归 -> 递归参数 & 终止条件 & 递推工作
func copyRandomList(head *Node) *Node {

	temp := make(map[*Node]*Node)
	return dfs(head, temp)
}

func dfs(head *Node, temp map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}
	if v, ok := temp[head]; ok {
		return v
	}
	copy := &Node{head.Val, nil, nil}
	temp[head] = copy
	copy.Next = dfs(head.Next, temp)
	copy.Random = dfs(head.Random, temp)
	return copy
}

func TestCopyRandomList(t *testing.T) {

}
