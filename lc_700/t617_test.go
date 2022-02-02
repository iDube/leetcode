package lc_700

// 617. 合并二叉树

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	tree := &TreeNode{}
	if root1 != nil && root2 != nil {
		tree.Val = root1.Val + root2.Val
		tree.Right = mergeTrees(root1.Right, root2.Right)
		tree.Left = mergeTrees(root1.Left, root2.Left)
	}
	if root1 != nil && root2 == nil {
		tree.Val = root1.Val
		tree.Right = mergeTrees(root1.Right, nil)
		tree.Left = mergeTrees(root1.Left, nil)
	}
	if root1 == nil && root2 != nil {
		tree.Val = root2.Val
		tree.Right = mergeTrees(nil, root2.Right)
		tree.Left = mergeTrees(nil, root2.Left)
	}
	return tree
}
