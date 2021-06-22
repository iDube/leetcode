package lc_6000

// TODO 5179. 将二叉搜索树变平衡

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {

	// 提取数据
	data := root.Data()

	// 构建 AVL 树
	root = BuildAvlTree(data)

	return root

}

func (this *TreeNode) Data() []int {
	return []int{}
}

func BuildAvlTree(data []int) *TreeNode {
	return &TreeNode{}
}
