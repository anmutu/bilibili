/*
  author='du'
  date='2020/9/3 21:19'
*/

// https://github.com/anmutu/bilibili
//# https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
//
//# 给定两个数组，编写一个函数来计算它们的交集.
//
//# example
//# 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//# 输出: [4,9]
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
//注意:输出结果中的每个元素一定是唯一的。
package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums1 []int = []int{4, 9, 5}
	var nums2 []int = []int{9, 4, 9, 8, 4}
	fmt.Println(interSectionOf2Arrays(nums1, nums2))

	var n1 []int = []int{1, 2, 2, 1}
	var n2 []int = []int{2, 2}
	fmt.Println(interSectionOf2Arrays(n1, n2))
}

//双指针解法(用排序和双指针解决）。
//1.先将两个数组排序。
//2.从两个数组的第一个元素开始比较。
//3.如果两个数相等，就将这个元素放入我们要定义的res里去。且将两个元素的下标都往后移一步。
//4.如果两个不相等，则将小的元素对应的下标往后移一下，接着比较。
//5.注意终止条件是要某个数组比较完毕了。
func interSectionOf2Arrays(nums1 []int, nums2 []int) []int {
	var res []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	p1 := 0
	p2 := 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			res = append(res, nums1[p1])
			p1 += 1
			p2 += 1
		} else if nums1[p1] < nums2[p2] {
			p1 += 1
		} else {
			p2 += 1
		}
	}
	return res
}
