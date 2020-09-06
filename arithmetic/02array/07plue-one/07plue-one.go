/*
  author='du'
  date='2020/9/3 21:19'
*/

//# https://leetcode-cn.com/problems/plus-one/
//
//# 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//# 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//# 你可以假设除了整数 0 之外，这个整数不会以零开头
//
//# example
//# 输入: [4,3,2,1]
//# 输出: [4,3,2,2]

package main

import "fmt"

func main() {
	var nums []int = []int{9, 9, 9, 9}
	fmt.Println(plusOne(nums))
	var nums1 []int = []int{8, 8, 8, 8}
	fmt.Println(plusOne(nums1))
}

func plusOne(digits []int) []int {
	l := len(digits)

	if l == 0 {
		return []int{1}
	}

	for i := l - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

//想一想有没有其他的解法呢。
