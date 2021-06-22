package ms_40

import (
	"fmt"
	"testing"
)

// 剑指 Offer 32 - III. 从上到下打印二叉树 III

// 和 32-2 相比，返回数组对用树中的节点，偶数层（root为第0层）从左到右，奇数层是从右到左

func levelOrder3(root *TreeNode) [][]int {

	data := [][]int{}
	if root == nil {
		return data
	}

	temp := []*TreeNode{}
	temp = append(temp, root)

	// deep 层
	deep := 0
	for len(temp) > 0 {
		t, d := ddd3(temp, deep)
		temp = t
		data = append(data, d)
		deep++
	}

	return data

}

func ddd3(temp []*TreeNode, deep int) ([]*TreeNode, []int) {
	d := []int{}

	// 每次处理一层，这层有 size 个节点
	for size := len(temp); size > 0; size-- {

		// 偶数层（root为第0层）从左到右，奇数层是从右到左
		if deep%2 == 0 {
			d = append(d, temp[0].Val)
		} else {
			d = append([]int{temp[0].Val}, d...)
		}

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

func TestLevelOrder3(t *testing.T) {
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

	fmt.Println(levelOrder3(root))
}
