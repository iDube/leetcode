package ms_40

import (
	"fmt"
	"testing"
)

// 面试题32 - II. 从上到下打印二叉树 II

// 先序遍历，带上层级信息

func levelOrder2(root *TreeNode) [][]int {

	data := [][]int{}
	if root == nil {
		return data
	}

	temp := []*TreeNode{}
	temp = append(temp, root)

	for len(temp) > 0 {
		t, d := ddd(temp)
		temp = t
		data = append(data, d)
	}

	return data

}

func ddd(temp []*TreeNode) ([]*TreeNode, []int) {
	d := []int{}

	// 每次处理一层，这层有 size 个节点
	for size := len(temp); size > 0; size-- {

		d = append(d, temp[0].Val)

		if temp[0].Left != nil {
			temp = append(temp, temp[0].Left)
		}
		if temp[0].Right != nil {
			temp = append(temp, temp[0].Right)
		}

		temp = temp[1:]
	}

	return temp, d
}

func TestLevelOrder2(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Right: &TreeNode{
				Val: 15,
			},
			Left: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(levelOrder2(root))
}
