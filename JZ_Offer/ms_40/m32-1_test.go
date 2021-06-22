package ms_40

import (
	"fmt"
	"testing"
)

// TODO 剑指 Offer 32 - I. 从上到下打印二叉树

// 加个 2 避免和 32-2 一样
func levelOrder(root *TreeNode) []int {

	result := []int{}
	if root == nil {
		return result
	}

	// 对节点进行缓存, 把 temp 当队列用
	temp := []*TreeNode{root}

	// 循环遍历
	for len(temp) > 0 {

		result = append(result, temp[0].Val)

		// 当队列用，往后追加
		if temp[0].Left != nil {
			temp = append(temp, temp[0].Left)
		}
		if temp[0].Right != nil {
			temp = append(temp, temp[0].Right)
		}

		// 当队列用，消费前面的
		temp = temp[1:]
	}

	return result

}

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 15,
			},
		},
	}

	fmt.Println(levelOrder(root))
}
