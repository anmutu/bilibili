/*
  author='du'
  date='2020/10/18 4:52'
*/
package main

import "fmt"

//给定一个整数数组 nums，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//示例:
//输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6

func main() {
	var nums []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//var nums []int=[]int{1,2,-20,3,4}
	s := 7
	fmt.Println(getSum(nums, s))
}

func getSum(nums []int, s int) int {
	res := nums[0]
	subRes := 0
	for i := 0; i < len(nums)-1; i++ {
		subRes += nums[i]
		if res < subRes {
			res = subRes
		}
		if subRes < 0 {
			subRes = 0
		}
	}
	return res
}
