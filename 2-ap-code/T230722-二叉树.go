package main

import "fmt"

/*
	二叉树遍历：
		前序遍历、中序遍历、后序遍历（按根节点的访问划分，左子树先于右子树）
		DFS深度搜索、BFS层次搜索
	易错点：
		1. 分清出root、node
		2. DFS：从上到下（传入指针）、从下到上（返回结果）
*/

// 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历：递归
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	//根-左-右
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 前序遍历：非递归，栈
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root) // 入栈
			root = root.Left
		}
		node := stack[len(stack)-1] // 出栈
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 中序遍历：非递归，栈
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root) // 进栈
			root = root.Left
		}
		node := stack[len(stack)-1] // 出栈
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

// 后序遍历：非递归，栈
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	lastVisted := &TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 先观察，先不急着出栈
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisted { //没有右节点，或者你的右节点刚访问过了，那就轮到你
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastVisted = node
		} else {
			root = node.Right
		}
	}
	return result
}

// DFS深度搜索：从上到下--就是前序遍历
func preorderTraversalDFS(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// DFS深度搜索-从下到上：分治法
func preorderTraversalDFSDiv(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	left := divideAndConquer(root.Left) // 分治
	right := divideAndConquer(root.Right)

	result = append(result, root.Val) //合并
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// BFS层次遍历-队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root) // 进队
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue) //记录当前层的个数
		for i := 0; i < l; i++ {
			node := queue[0] // 出队
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node1.Left = node2
	node1.Right = node3
	node3.Left = node4

	//preorderTraversal(node1)
	//result := preorderTraversal2(node1)
	//result := inorderTraversal(node1)
	//result := postorderTraversal(node1)
	//result := preorderTraversalDFS(node1)
	//result := preorderTraversalDFSDiv(node1)
	result := levelOrder(node1)
	fmt.Println(result)
}
