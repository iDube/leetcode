package ms_40

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 剑指 Offer 33. 二叉搜索树的后序遍历序列

// 判断一个序列是不是二叉搜索树的后续遍历
// 二叉排序树的中序遍历是个有序序列
// 找二叉排序树后序遍历序列的特点：
// 最后一个是根节点

func verifyPostorder(postorder []int) bool {
	if len(postorder) < 2 {
		return true
	}

	root := postorder[len(postorder)-1]

	// 根节点的值将序列分乘前后两个，都大于根节点的是右子树，都小于根节点的序列是左子树

	// 从后向前寻找分割点
	cr := len(postorder) - 2
	for postorder[cr] > root && cr >= 1 {
		cr--
	}

	// 从前向后寻找分割点
	cl := 0
	for postorder[cl] < root && cl <= len(postorder)-2 {
		cl++
	}

	if cl < cr {
		return false
	}

	// 继续递归
	la := postorder[:cr+1]
	ra := postorder[cr+1 : len(postorder)-1]
	return verifyPostorder(la) && verifyPostorder(ra)
}

func TestVerifyPostorder(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(verifyPostorder([]int{5, 10, 6, 9, 4}), false)
	assertions.Equal(verifyPostorder([]int{1, 2, 5, 10, 6, 9, 4, 3}), false)
	assertions.Equal(verifyPostorder([]int{1, 3, 2, 6, 5}), true)
	assertions.Equal(verifyPostorder([]int{1, 6, 3, 2, 5}), false)
}
