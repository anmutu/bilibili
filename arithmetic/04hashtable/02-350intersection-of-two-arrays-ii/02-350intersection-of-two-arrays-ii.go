/*
  author='du'
  date='2021/2/9 21:18'
  github='https://github.com/anmutu'
*/

// https://github.com/anmutu/bilibili
//# 给定两个数组，编写一个函数来计算它们的交集.
//# example
//# 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//# 输出: [4,9]
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]

package main

import "fmt"

func main() {
	var n1 []int = []int{1, 2, 2, 1}
	var n2 []int = []int{2, 2}
	fmt.Println(interSectionOf2ArraysMap(n1, n2))
}

//想办法构造一相查找表。然后再这个查找表里去找元素是否存在。
//n1=>map[1:2 2:2]
//查找表：遍历一个数组，以其值为key，其数量为value。
//遍历另外一个数组，判断是否存在。
func interSectionOf2ArraysMap(nums1 []int, nums2 []int) (res []int) {
	hash := make(map[int]int)
	for _, v := range nums1 {
		hash[v] += 1
	}
	fmt.Println(hash)
	for _, v := range nums2 {
		if hash[v] > 0 {
			res = append(res, v)
			hash[v]--
		}
	}
	return res
}
