package ms_60

// 面试题55 - II. 平衡二叉树

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	depLeft := maxDepth(root.Left)
	depRight := maxDepth(root.Right)
	if depLeft-depRight == 1 || depLeft-depRight == -1 || depLeft == depRight {
		return isBalanced(root.Left) && isBalanced(root.Right)
	} else {
		return false
	}

}
