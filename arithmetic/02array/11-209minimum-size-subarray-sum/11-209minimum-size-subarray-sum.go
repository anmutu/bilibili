/*
  author='du'
  date='2020/10/18 22:14'
*/
package main

import (
	"fmt"
	"math"
)

//输入：s = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。

func main() {
	s := 7
	var nums []int = []int{2, 3, 1, 2, 4, 3}
	//var nums []int=[]int{1,2,7}
	fmt.Println(minSubArrayLen(s, nums))
}

//滑动窗口。
//会有一个不确定大小的窗口。
//定义一个左右指针left和right。
//一直会往右走。
//如果窗口里面的值大于【等于】目标值s的。则把左边的值减去，然后左指针往右走一步。
//如果窗口里面的值小于目标值s。则把右指针往右走一步，且将对应的值加入到窗口中去。
//会有多个窗口，所以在遍历后是需要去找出符合条件的值中的最小值的。
//条件：left<len(nums)
func minSubArrayLen(s int, nums []int) int {
	left, right := 0, -1
	sum := 0
	res := len(nums) + 1
	for left < len(nums) {
		if right+1 < len(nums) && sum < s {
			right++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		}
		if sum >= s {
			res = int(math.Min(float64(res), float64(right-left+1)))
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
