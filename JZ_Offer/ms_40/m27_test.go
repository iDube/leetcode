package ms_40

// 面试题27. 二叉树的镜像

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func mirrorTree(root *TreeNode) *TreeNode {

	if root != nil {
		tr := root.Left
		root.Left = mirrorTree(root.Right)
		root.Right = mirrorTree(tr)
	}
	return root

}
