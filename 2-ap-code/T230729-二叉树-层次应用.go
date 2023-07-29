package main

/*
	BFS层次应用

*/

// 102. 二叉树的层次遍历
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue) // 记录当前层的节点数
		for i := 0; i < l; i++ {
			node := queue[0] // 出队
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left) // 入队
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

// 107. 二叉树的层次遍历2
func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder2(root)
	reverse(result)
	return result
}
func reverse(nums [][]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 103. 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	flag := false
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0) //注意，指定size会自动填充默认零值，后续不能append，实在想提前分配空间，后续用下标赋值
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if flag {
			reverse2(tmp)
		}
		flag = !flag
		result = append(result, tmp)
	}
	return result
}

func reverse2(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
