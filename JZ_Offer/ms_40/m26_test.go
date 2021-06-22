package ms_40

import (
	"fmt"
	"testing"
)

// 剑指 Offer 26. 树的子结构

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	if A.Val != B.Val {
		return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
	} else {
		// 可能存在相同的节点，
		// 所以，A,B 值相同的话，用 isSubStructure2 向下递归，如果为子节点一样，一定是子树
		// 如果子节点不一样，继续把B去和A的子树进行比较
		if isSubStructure2(A.Left, B.Left) && isSubStructure2(A.Right, B.Right) {
			return true
		} else {
			return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
		}
	}

}

func isSubStructure2(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	} else {
		return isSubStructure2(A.Left, B.Left) && isSubStructure2(A.Right, B.Right)
	}

}

// 判断两个树是否相同
func isSame(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	}

	return isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
}

func TestIsSame(t *testing.T) {
	root1 := &TreeNode{
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

	root2 := &TreeNode{
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

	fmt.Println(isSame(root1, root2))
}

func TestIsSubStructure(t *testing.T) {
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	fmt.Println(isSubStructure(root1, root2))
}
