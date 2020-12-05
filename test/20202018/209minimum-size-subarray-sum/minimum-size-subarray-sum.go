/*
  author='du'
  date='2020/10/18 6:13'
*/
package main

import (
	"fmt"
	"math"
)

//给定一个含有n个正整数的数组和一个正整数s ，找出该数组中满足其和 ≥ s 的长度最小的
//连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
//
//示例：
//输入：s = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组[4,3]是该条件下的长度最小的子数组。

func main() {
	var nums []int = []int{2, 3, 1, 2, 4, 3}
	s := 7
	fmt.Println(getLength(nums, s))
}

//滑动窗口。需要两个
//右指针右移条件，1是sum<s,2是right+1<len(nums)
//right定义为-1是希望最开始的时候不包含任务元素。
//定义左右两个指针，left,right。
//当指针内的数相加和小于目标值s时，则右指针向右移动。且将此时的值加入到总和中。
//当指针内的数相加和大于等于目标值s时，则将最左的值从总和中减去，且把左指针往右移动。
//遍历条件，left<len(nums)
//另外在遍历后有符合题目要求的，则用Min函数去找到最小的数值。
func getLength(nums []int, s int) int {
	left, right, sum := 0, -1, 0
	res := len(nums) + 1
	for left < len(nums) {
		if right+1 < len(nums) && sum < s {
			right++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		}
		if sum >= s {
			res = int(math.Min(float64(res), float64(right-left+1)))
			fmt.Printf("%v=>%d,\n", nums[left:right+1], sum)
		}
	}
	return res
}
