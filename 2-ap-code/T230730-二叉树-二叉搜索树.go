package main

/*
	二叉搜索树的应用
*/

// 98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return traverse(root, nil, nil)
}
func traverse(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	val := root.Val
	if min != nil && val <= min.Val {
		return false
	}
	if max != nil && val >= max.Val {
		return false
	}

	left := traverse(root.Left, min, root)
	right := traverse(root.Right, root, max)
	return left && right
}

// 701.二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 1.返回条件
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	// 2. 分段处理
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
