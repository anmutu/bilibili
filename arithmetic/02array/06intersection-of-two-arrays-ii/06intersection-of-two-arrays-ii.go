/*
  author='du'
  date='2020/9/3 21:19'
*/

//# https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
//
//# 给定两个数组，编写一个函数来计算它们的交集.
//
//# example
//# 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//# 输出: [4,9]

package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums1 []int = []int{4, 9, 5}
	var nums2 []int = []int{9, 4, 9, 8, 4}
	res := interSectionOf2Arrays(nums1, nums2)
	fmt.Println(res)
}

//双指针解法
//1.先将两个数组排序。
//2.从两个数组的第一个元素开始比较。
//3.如果两个数相等，就将这个元素放入我们要定义的res里去。且将两个元素的下标都往后移一步。
//4.如果两个不相等，则将小的元素对应的下标往后移一下，接着比较。
//5.注意终止条件是要两个数组都走到他们对应数组的最后一个位置。
func interSectionOf2Arrays(nums1 []int, nums2 []int) []int {
	var res []int = []int{}
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

//另外，可以考虑下，如果一个nums很短，另外一个nums特别长，应该怎么处理会比较好一点。
