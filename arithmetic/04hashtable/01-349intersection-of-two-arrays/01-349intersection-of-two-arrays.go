/*
  author='du'
  date='2021/2/9 17:32'
  github='https://github.com/anmutu'
*/

//给定两个数组，编写一个函数来计算它们的交集。
//url:https://leetcode-cn.com/problems/intersection-of-two-arrays
//示例 1:
//输入:nums1 = [1,2,2,1], nums2 = [2,2]
//输出:[2]
//示例 2:
//输入:nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出:[9,4]
//注意:输出结果中的每个元素一定是唯一的。

package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums1 []int = []int{1, 2, 2, 1}
	var nums2 []int = []int{2, 2}
	fmt.Println(interSectionOf2Arrays(nums1, nums2))
	var n1 []int = []int{1, 2, 2, 1}
	var n2 []int = []int{2, 2}
	fmt.Println(interSectionOf2ArraysMap(n1, n2))
}

//排序和双指针。
//与349相比，就是增加了一个判断。
//在把元素加入之前，判断这个元素是否大于结果集中的最后一个元素。
//1.大于才把这个值加入到结果集中。
//2.大小于（等于）了依旧是需要把两个指针都分别往后移的。
//349讲解视频：https://www.bilibili.com/video/BV1E54y1y748
func interSectionOf2Arrays(nums1 []int, nums2 []int) []int {
	var res []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	p1 := 0
	p2 := 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			//这里和349相比就是要判断一下，当前要加入的这个数是否大于结果集中最后的一上元素。
			if res == nil || nums1[p1] > res[len(res)-1] {
				res = append(res, nums1[p1])
			}
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

//定义一个map去收集nums1里的元素。
//存在的话，则以这个数据对应元素的值当map的key，给其对应的value为true。这里是需要遍历一次的。
//遍历另外一个数组，如果map以nums对应元素为key是为true的话，那么就把这个值加入到我们的结果集当中。
//注意：加入后还需要把这个对应的map改为false。
func interSectionOf2ArraysMap(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	hash := make(map[int]bool)
	for _, v := range nums1 {
		hash[v] = true
	}
	for _, v := range nums2 {
		if hash[v] == true {
			res = append(res, v)
			hash[v] = false
		}
	}
	return res
}
