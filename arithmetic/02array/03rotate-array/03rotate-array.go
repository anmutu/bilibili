/*
  author='du'
  date='2020/9/3 21:18'
*/

//# 旋转数组
//# https://leetcode-cn.com/problems/rotate-array/
//# 给定一个数组，将数组中的元素向右移动k个位置，其中k是非负数。
//
//# example
//# 输入: [1,2,3,4,5,6,7] 和 k = 3
//# 输出: [5,6,7,1,2,3,4]
//# 解释:
//# 向右旋转 1 步: [7,1,2,3,4,5,6]
//# 向右旋转 2 步: [6,7,1,2,3,4,5]
//# 向右旋转 3 步: [5,6,7,1,2,3,4]

package main

import (
	"fmt"
)

func main() {
	var nums []int = []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	res := rotateArray(nums, k)
	fmt.Println(res)
}

func rotateArray(nums []int, k int) []int {
	temp := nums[len(nums)-k:]
	nums = nums[0 : len(nums)-k]
	nums = append(temp, nums...)
	return nums
}
