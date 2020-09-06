/*
  author='du'
  date='2020/9/3 21:19'
*/

//# https://leetcode-cn.com/problems/move-zeroes/
//
//# 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//# example
//# 输入: [0,1,0,3,12]
//# 输出: [1,3,12,0,0]

package main

import "fmt"

func main() {
	var nums []int = []int{0, 1, 0, 3, 12}
	res := moveZeros(nums)
	fmt.Println(res)

}

//删除追加法。注意golang相关的写法。
func moveZeros(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}
	return nums
}

//func moveZeros(nums []int){
//	left:=0
//	right:=0
//	for right<len(nums){
//		if nums[left]==0{
//			nums[len(nums)]=0
//			nums[left:len(nums)-1]=nums[left+1:len(nums)]
//			right+=1
//		}else{
//			left+=1
//			right+=1
//		}
//	}
//
//}
