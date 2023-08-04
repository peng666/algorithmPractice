package main

import "strconv"

/*
	栈：DFS
	队列：BFS

	155. 最小栈
	150. 逆波兰表达式求值
	394. 字符串解码
	94. 二叉树的中序遍历
	133. 克隆图
	200. 岛屿数量
*/

// 155. 最小栈---两个栈实现
type MinStack struct {
	min   []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}
func (m *MinStack) Push(x int) {
	min := m.GetMin()
	if x < min {
		m.min = append(m.min, x)
	} else {
		m.min = append(m.min, min)
	}
	m.stack = append(m.stack, x)
}

func (m *MinStack) Pop() {
	if len(m.stack) == 0 {
		return
	}
	m.stack = m.stack[:len(m.stack)-1]
	m.min = m.min[:len(m.min)-1]
}

func (m *MinStack) Top() int {
	if len(m.stack) == 0 {
		return 0
	}
	return m.stack[len(m.stack)-1]
}
func (m *MinStack) GetMin() int {
	if len(m.min) == 0 {
		return 1 << 31
	}
	min := m.min[len(m.min)-1]
	return min
}

// 150. 逆波兰表达式求值
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return -1
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result int
			switch tokens[i] {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			stack = append(stack, result)
		default:
			val, _ := strconv.Atoi(tokens[i])
			stack = append(stack, val)
		}
	}
	return stack[0]
}

// 394. 字符串解码
func decodeString(s string) string {
	if len(s) == 0 {
		return s
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			tmp := make([]byte, 0) // 存放本次要打印的字符串
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				tmp = append(tmp, v)
			}
			stack = stack[:len(stack)-1] // 将'['弹出去
			idx := 1
			for len(stack) >= idx && stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' { // 获取数字位数
				idx++
			}
			num := stack[len(stack)-idx+1:] // 数字
			stack = stack[:len(stack)-idx+1]
			count, _ := strconv.Atoi(string(num))
			for j := 0; j < count; j++ {
				for k := len(tmp) - 1; k >= 0; k-- { // 倒序遍历
					stack = append(stack, tmp[k]) // 将解密后的字符串加回在stack中
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

// 94. 二叉树的中序遍历
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

type Node2 struct {
	Val       int
	Neighbors []*Node2
}

// 133. 克隆图
func cloneGraph(node *Node2) *Node2 {
	visited := make(map[*Node2]*Node2)
	return clone(node, visited)
}
func clone(node *Node2, visited map[*Node2]*Node2) *Node2 {
	if node == nil {
		return nil
	}

	if v, ok := visited[node]; ok {
		return v
	}

	newNode := &Node2{
		Val:       node.Val,
		Neighbors: make([]*Node2, len(node.Neighbors)),
	}
	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode
}

// 200.岛屿数量
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' { // 每遇到一个岛屿，就把它全部标记为0
				dfs2(grid, i, j)
				count++
			}
		}
	}
	return count
}
func dfs2(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		dfs2(grid, i-1, j)
		dfs2(grid, i+1, j)
		dfs2(grid, i, j-1)
		dfs2(grid, i, j+1)
	}
	return
}
