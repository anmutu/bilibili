/*
  author='du'
  date='2020/11/1 15:28'
*/
package main

//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//满足要求的三元组集合为：
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

func main() {
	//var nums []int=[]int{-1, 0, 1, 2, -1, -4}
	//fmt.Println(threeSum(nums))
}

//-1，需要后面 0, 1, 2, -1, -4找出两个数相加等于-(-1)
//先遍历这个数组，以第一个数为target,找出剩下数组里面的数使之满足b+c=-target。
//遍历的条件:len(nums)-2
//然后里面的就是two-sum的逻辑了。
//需要去重的逻辑。
//1.target的去重。
//2.left和right的去重。
//func threeSum(nums []int) [][]int {
//  var res [][]=[][]int{}
//  for i:=0;i<len(nums)-2;i++{
//  	target:=nums[i]
//  	left,right:=i+1,len()
//  }
//  return res
//}
