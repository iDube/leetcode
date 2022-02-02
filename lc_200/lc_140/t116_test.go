package lc_140

import (
	"fmt"
	"testing"
)

// 116. 填充每个节点的下一个右侧节点指针

// 广度优先遍历
func connect(root *Node) *Node {

	// 题目给出，root是完美二叉树

	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {

		// 每次取树的一层
		qTemp := queue
		queue = []*Node{}

		length := len(qTemp)
		for i := 0; i < length; i++ {
			if i == length-1 {
				// 最后一个
				qTemp[i].Next = nil
			} else {
				qTemp[i].Next = qTemp[i+1]
			}
			if qTemp[i].Left != nil {
				queue = append(queue, qTemp[i].Left)
			}
			if qTemp[i].Right != nil {
				queue = append(queue, qTemp[i].Right)
			}
		}

	}

	return root
}

func Test_connect(t *testing.T) {

	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}
	fmt.Println(root)
	fmt.Println(connect(root))

}
