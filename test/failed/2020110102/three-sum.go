/*
  author='du'
  date='2020/11/1 15:39'
*/
package main

import (
	"fmt"
	"sort"
)

//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//满足要求的三元组集合为：
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

//1.最先是要排序
//2.想办法转换我成two-sum
//比如，现在是以第0个数，也就是-1做为target了
//那么这就转化成了求-1后面的数是不是有两个数相加等于-1的相反数。
//条件：len(nums)-2
//3.去重的逻辑。
//3a.target的去重 。
//3b.left和right的去重 。
func main() {
	var nums []int = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var res [][]int = [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == target {
				r := []int{nums[i], nums[left], nums[right]}
				res = append(res, r)
				left++
				right--
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
