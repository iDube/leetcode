package lc_560

import (
	"testing"
)

// 543. 二叉树的直径

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 所有左右子数高度和最大值
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hightOfBinaryTree(root)
	return maxHight
}

var maxHight int = 0

// 求树的高度, 为了减少遍历次数，同时记录记录左右子树高度和的最大值
func hightOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := hightOfBinaryTree(root.Left)
	rh := hightOfBinaryTree(root.Right)
	maxHight = max(lh+rh, maxHight) // 遍历时，记录左右子树高度和的最大值
	return max(lh, rh) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Test_hightOfBinaryTree(t *testing.T) {
	// TODO
}

func Test_diameterOfBinaryTree(t *testing.T) {
	// TODO
	// 这个错了
	// [4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]
}

// 将数组转成一棵完全二叉树, 按行平铺
// 如 [1,2,3,4,5,6]
//			1
//		   / \
//        2   3
//       / \ /
//      4  56
func arrayToTree(array []int) *TreeNode {
	// TODO
	return nil
}
