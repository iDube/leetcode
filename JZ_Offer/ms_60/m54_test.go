package ms_60

import "testing"

// 面试题54. 二叉搜索树的第k大节点

// 中序遍历，找到第k个
func kthLargest(root *TreeNode, k int) int {
	temp := mid(root)
	// 题目给出，树内至少一个元素; 1 ≤ k ≤ 二叉搜索树元素个数
	return temp[len(temp)-k]
}

func mid(root *TreeNode) []int {
	array := []int{}
	if root == nil {
		return array
	}

	array = append(array, mid(root.Left)...)
	array = append(array, root.Val)
	array = append(array, mid(root.Right)...)
	return array
}

func TestKthLargest(t *testing.T) {

	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}

	println(kthLargest(root, 2))
	println(kthLargest(root, 5))

}
