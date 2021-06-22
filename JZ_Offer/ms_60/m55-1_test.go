package ms_60

import "testing"

// 面试题55 - I. 二叉树的深度

func maxDepth(root *TreeNode) int {
	return depth(root, 0)
}

func depth(root *TreeNode, dep int) int {
	if root == nil {
		return dep
	}

	dep++
	leftDep := depth(root.Left, dep)
	rightDep := depth(root.Right, dep)

	if dep < leftDep {
		dep = leftDep
	}
	if dep < rightDep {
		dep = rightDep
	}
	return dep
}

func TestMaxDepth(t *testing.T) {
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
	println(maxDepth(root))
}
