/*
  author='du'
  date='2020/11/1 17:05'
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

//a+b+c=0
//0-a=b+c
//主体思想：不用暴力破解。时间复杂度太大。想办法转换成two-sum的题。
//先排序。
//-1 =>target=0-(-1)=1  [0,1,2, -1, -4]去找是否有两个数相加等于target。
//条件：len(nums)-2
//去重。
//1.target的一个去重。
//2.left和一个right的去重。
//[-1,0,0,0,1,1,1]
func main() {
	var nums []int = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var res [][]int = [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			return nil
		}
		target := 0 - nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == target {
				r := []int{nums[i], nums[left], nums[right]}
				res = append(res, r)
				right--
				for nums[right] == nums[right+1] {
					right--
				}
				left++
				for nums[left] == nums[left-1] {
					left++
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
