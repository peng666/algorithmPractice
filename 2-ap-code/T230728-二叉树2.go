package main

import (
	"fmt"
	"math"
)

/*
	分治法应用
		1. 递归返回条件
		2. 分段处理
		3. 合并结果
*/

// MergeSort 归并排序
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}
func mergeSort(nums []int) []int {
	// 1.返回条件
	if len(nums) <= 1 {
		return nums
	}

	// 2.分段处理
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	// 3.合并结果
	result := merge(left, right)
	return result
}
func merge(left, right []int) []int {
	l := 0
	r := 0
	result := []int{}
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

// QuickSort 快速排序--类似分治，但没有合并过程（快排是原地交换）
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, start, end int) {
	if start < end { // 1.返回条件
		pivot := partition(nums, start, end) // 2.分段处理
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}
func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p { // 跟基准元素比较，小的放在左边
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end) // 最后把基准值放到分界线中间的位置，分隔开（左边小，右边大）
	return i           // 返回的是分隔线下标
}
func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

// 经典题目应用-分治法
// 104.二叉树的最大深度
func maxDepth(root *TreeNode) int {
	// 1.结束条件
	if root == nil {
		return 0
	}
	// 2.分段处理
	depth := 1
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	// 3.合并结果
	if left > right {
		depth = depth + left
	} else {
		depth = depth + right
	}
	return depth
}

// 110.平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, flag := maxDepthBala(root)
	return flag
}
func maxDepthBala(root *TreeNode) (int, bool) {
	// 1.返回条件
	if root == nil {
		return 0, true
	}
	// 2.分段处理
	depthL, flagL := maxDepthBala(root.Left)
	depthR, flagR := maxDepthBala(root.Right)

	// 3.合并结果
	depth := depthL
	if depthR > depthL {
		depth = depthR
	}
	flag := flagL && flagR && abs(depthL, depthR) <= 1
	return depth + 1, flag // 计算数的高度记得要+1当前层
}

func abs(x, y int) int {
	if x >= y {
		return x - y
	}
	return y - x
}

// 124.二叉树的最大路径和
type ResultType struct {
	SinglePath int
	MaxPath    int
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}
func helper(root *TreeNode) *ResultType {
	// 1.返回条件
	if root == nil {
		return &ResultType{SinglePath: 0, MaxPath: math.MinInt32}
	}
	// 2.分段处理
	left := helper(root.Left)
	right := helper(root.Right)

	// 3.合并结果
	result := &ResultType{}
	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}
	maxPath := max(right.MaxPath, left.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 4, 3, 78, 4, 55, 12}
	//result := MergeSort(nums)
	result := QuickSort(nums)

	fmt.Println(result)
}
