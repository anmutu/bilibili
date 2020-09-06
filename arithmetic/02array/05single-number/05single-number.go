/*
  author='du'
  date='2020/9/3 21:18'
*/
package main

import (
	"fmt"
	"sort"
)

//# https://leetcode-cn.com/problems/single-number/
//
//# 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//# 说明：
//# 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//# example
//# 输入: [4,1,2,1,2]
//# 输出: 4

func main() {
	var nums []int = []int{4, 1, 2, 1, 2}
	res := singleNumber(nums)
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	i := 0
	for i < len(nums)-2 {
		if nums[i] == nums[i+1] {
			i += 2
		} else {
			return nums[i]
		}
	}
	return nums[len(nums)-1] //注意这里要-1.
}
