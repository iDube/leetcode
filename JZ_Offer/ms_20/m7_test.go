package ms_20

import (
	"testing"
)

// 面试题07. 重建二叉树
// 105. 从前序与中序遍历序列构造二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 看了书里的解析，
// 1. 前序遍历序列的第一个数字1就是根结点的值。
// 2. 再中序遍历中，从头往后找到根节点，根节点前面部分就是左子树的中序遍历，后面部分是右子树的中序遍历
// 3. 从2中提取左右子树数组长度，在前序遍历中，也可标出左子树的前序遍历和右子树的前序遍历
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 边界值检查
	if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}

	l := len(preorder)

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			root.Left = buildTree(preorder[1:i+1], inorder[0:i])
			root.Right = buildTree(preorder[i+1:l], inorder[i+1:l])
		}
	}

	return root
}

func Test_buildTree(t *testing.T) {

}
