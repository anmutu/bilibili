/*
  author='du'
  date='2020/11/1 2:36'
*/
package main

import (
	"fmt"
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
//使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//注意：答案中不可以包含重复的三元组。
//示例：
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//满足要求的三元组集合为：
//[
//[-1, 0, 1],
//[-1, -1, 2]

func main() {
	var nums []int = []int{-1, 0, 1, 2, -1, -4}
	var nums1 []int = []int{-1, -1, -1, 0, 1}
	var nums2 []int =  []int{-2, -1, 1, 1, 1, 1, 1, 1}
	var nums3 []int = []int{-2, -1,-1, 1, 1}
	fmt.Println(threeSum(nums))
	fmt.Println(threeSum(nums1))
	fmt.Println(threeSum(nums2))
	fmt.Println(threeSum(nums3))
}

func threeSum(nums []int) [][]int {
	var res [][]int = [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		//if i>0&&nums[i]==nums[i-1]{
		//	continue
		//}
		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == target {
				r := []int{nums[i], nums[left], nums[right]}
				res = append(res, r)
				left++
				//for left < right && nums[left] == nums[left-1] {
				//	left++
				//}
				right--
				//for left < right && nums[right] == nums[right+1] {
				//	right--
				//}
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
