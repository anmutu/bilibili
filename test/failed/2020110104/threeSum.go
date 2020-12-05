/*
  author='du'
  date='2020/11/1 16:48'
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

//主体思想：把这个three-sum的题转换成two-sum的题去做。
//a+b+c=0
//0-a=b+c
//排序。
//遍历数组，以0-nums[i]做为target。
//再剩下的数里面去找是否有两个数相加等于target。
//至此，就是一个two-sum的一个逻辑了。
//去重。
//1.target的去重。
//2.left和right一个去重 。

func main() {
	var nums []int=[]int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
   var res [][]int=[][]int{}
   sort.Ints(nums)
   for i:=0;i<len(nums)-2;i++{
   	target:=0-nums[i]
   	left,right:=i+1,len(nums)-1
   	for left<right{
   		if nums[left]+nums[right]==target{
   			r:=[]int{nums[i],nums[left],nums[right]}
   			res=append(res,r)
   			right--
   			left++
		}else if nums[left]+nums[right]>target{
			right--
		}else{
			left++
		}
	}
   }
   return res
}