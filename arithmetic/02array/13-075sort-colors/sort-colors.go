/*
  author='du'
  date='2021/1/16 12:35'
  github='https://github.com/anmutu'
*/
package main

import "fmt"

//链接：https://leetcode-cn.com/problems/sort-colors
//给定一个包含红色、白色和蓝色，一共n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，
//并按照红色、白色、蓝色顺序排列。
//此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
//注意：请不要使用代码库中的排序函数来解决这道题。
//进阶：
//你能想出一个仅使用常数空间的一趟扫描算法吗？
//示例 1：
//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]

func main() {
	var nums []int = []int{2, 0, 2, 1, 1, 0}
	setColors(nums)
	fmt.Println(nums)
}

func setColors(nums []int) {
	left, right := -1, len(nums)
	for i := 0; i < right; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 0 {
			left++
			nums[left], nums[i] = nums[i], nums[left]
			i++
		} else {
			right--
			nums[right], nums[i] = nums[i], nums[right]
		}
	}
}
