/*
  author='du'
  date='2020/9/3 21:18'
*/

//# https://leetcode-cn.com/problems/contains-duplicate/
//
//# 给定一个整数数组，判断是否存在重复元素。
//# 如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
//
//# example
//# 输入: [1,2,3,1]
//# 输出: true

package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums []int = []int{1, 2, 3, 1}
	res := containsDuplicate(nums)
	fmt.Println(res)
}

//存在重复元素返回true
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	i := 0
	for i < len(nums) {
		if nums[i] == nums[i+1] {
			return true
		} else {
			i += 2
		}
	}
	return false
}
