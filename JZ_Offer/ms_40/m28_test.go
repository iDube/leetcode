package ms_40

import "fmt"

// 面试题28. 对称的二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思考：
// 1. 能不能通过遍历来做？通过先中后序遍历得到的数组里找到些许规律
//   还真找到规律，先序遍历反过来刚好是后续遍历, 但是空节点时会出问题，要标记下空节点，用

// 执行超时了，可能是因为下面那个循环
func isSymmetric(root *TreeNode) bool {
	rf := RootFirst(root)
	fmt.Println(rf)
	rl := RootList(root)
	fmt.Println(rl)

	for i, j := 0, len(rl)-1; i < len(rf) && j >= 0; {
		if rf[i] != rl[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 先序遍历
func RootFirst(root *TreeNode) []int {
	if root == nil {
		return []int{-1 << 63}
	}
	array := []int{}
	array = append(array, root.Val)
	array = append(array, RootFirst(root.Left)...)
	array = append(array, RootFirst(root.Right)...)
	return array
}

// 后续遍历
func RootList(root *TreeNode) []int {
	if root == nil {
		return []int{-1 << 63}
	}
	array := []int{}
	array = append(array, RootList(root.Left)...)
	array = append(array, RootList(root.Right)...)
	array = append(array, root.Val)
	return array
}
