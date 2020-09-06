/*
  author='du'
  date='2020/9/3 21:20'
*/

//# https://leetcode-cn.com/problems/two-sum/
//
//# 给定一个整数数组 nums和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
//# 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//# example
//# 给定 nums = [2, 7, 11, 15], target = 9
//# 因为 nums[0] + nums[1] = 2 + 7 = 9
//# 所以返回 [0, 1]

package main

import "fmt"

func main() {
	var nums []int = []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}

//暴力解法
//func twoSum(nums []int,target int)[]int{
//	for i:=0;i<len(nums);i++{
//		for j:=i+1;j<len(nums);j++{
//			if nums[i]+nums[j]==target{
//				return []int{i,j}
//			}
//		}
//	}
//	return nil
//}

//对撞指针
func twoSum(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l]+nums[r] == target {
			return []int{l, r}
		} else if nums[l]+nums[r] > target {
			r--
		} else {
			l++
		}
	}
	return []int{l, r}
}

//func twoSum(nums []int, target int) []int {
//	// 申请map
//	m := make(map[int]int)
//	// 遍历
//	for index, num := range nums {
//		// 到map中去查找目标2
//		key, ok := m[target - num]
//		// 找不到直接保存目标2
//		if !ok{
//			m[num] = index
//		}else {
//			// 找到直接返回目标1、目标2的下标
//			return []int{index, key}
//		}
//	}
//	return nil
//}
//
