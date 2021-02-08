/*
  author='du'
  date='2020/9/3 21:20'
*/

//# https://leetcode-cn.com/problems/two-sum/
//  https://github.com/anmutu/bilibili/blob/master/arithmetic/02array/09two-sum/09two-sum.go
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
	var nums []int = []int{2, 11, 15, 7}
	target := 9
	res := twoSum1(nums, target)
	fmt.Println(res)
	fmt.Println(twoSum2(nums, target))
	fmt.Println(twoSum3(nums, target))

}

//暴力解法
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//对撞指针
func twoSum2(nums []int, target int) []int {
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

//map
//nums = [2, 7, 11, 15]
//数组
//value 下标
//  2   0
//  7   1
//  11  2
//  15  3
//想办法把这个转换成map。然后再map里面去寻找。
//map
//key value
// 2   0
// 7   1
// 11  2
// 15 3
func twoSum3(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for index, value := range nums {
		if p, ok := hashTable[target-value]; ok {
			return []int{p, index}
		}
		hashTable[value] = index
	}
	//for k,v :=range hashTable{
	//	if p,ok:=hashTable[target-k];ok{
	//		return []int{p,v}
	//	}
	//}
	return nil
}
