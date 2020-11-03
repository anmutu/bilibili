/*
  author='du'
  date='2020/11/2 2:18'
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

//a+b+c+0
//0-a=b+c
//[-4,-1,-1,0,1,2]
//target=0-(-4)=4  [-1,-1,0,1,2] 是否有两数相加等于4。
//target=0-(-1)=1  [-1,0,1,2]是否有两数相加等于1。
//主体思想：不用暴力的解法 。想办法 把这道题转换成two-sum的题去做。
//第一步，排序，按从小到大排序。
//取第i个元素，以0-nums[i]为target,在后面的数里面去找看是否有两个数相加等于target。 【已更正视频中i为nums[i]】
//条件：len(nums)-2
//去重。
//1.target的去重。
//2.left和right的去重。
func main() {
	//var nums []int = []int{-1, 0, 1, 2, -1, -4}
	var nums []int = []int{-1, 0, 0, 0, 0, 1, 1, 1, 1}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var res [][]int = [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[0] > 0 {
			return nil
		}
		target := 0 - nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == target {
				r := []int{nums[i], nums[left], nums[right]}
				res = append(res, r)
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
