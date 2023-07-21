package main

import "fmt"

/*
	回溯穷举法
*/

func subset(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

func backtrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}

func main() {
	nums := []int{2, 5}
	ret := subset(nums)
	fmt.Println(ret)
}
