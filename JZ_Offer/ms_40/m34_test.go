package ms_40

import "testing"

// 剑指 Offer 34. 二叉树中和为某一值的路径

// 这个路径一定要到叶子节点
// 节点有负数

// BFS ? 递归
func pathSum(root *TreeNode, sum int) [][]int {

	if root == nil {
		return [][]int{}
	}

	// root 的值等于sum, 且已经是 叶子节点。一定要到叶子节点，因为节点可能是负值
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	// 向下递归
	lr := pathSum(root.Left, sum-root.Val)
	rr := pathSum(root.Right, sum-root.Val)

	// root 节点值追加到数组前部
	r := append(lr, rr...)
	for i := range r {
		r[i] = append([]int{root.Val}, r[i]...)
	}
	return r
}

func TestPathSum(t *testing.T) {

}
