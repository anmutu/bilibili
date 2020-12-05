/*
  author='du'
  date='2020/11/1 16:16'
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
//b+c=0-a
//主体思想：把这个three-sum转换成two-sum。
//先对数组进行排序。
//遍历数组，选取0-nums[i]为target。
//然后再剩下的数里面去找数，看是否有两个数相加等于这个target。
//去重。
//1.target的去重。
//2.left和right的去重 。

func main() {
	//var nums []int=[]int{-1, 0, 1, 2, -1, -4}
	//fmt.Println(threeSum(nums))
	var nums1 []int=[]int{-1,0,1}
	fmt.Println(threeSum(nums1))
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