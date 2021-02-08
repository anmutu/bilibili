package main

import (
	"fmt"
	"sort"
)

// 给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
// 输入：nums = [3,0,1]
// 输出：2
// 解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。

// 输入：nums = [9,6,4,2,3,5,7,0,1]
// 输出：8
// 解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。

func main() {
	arr := []int{3, 0, 1}
	fmt.Println(missingNumber(arr))
}

//第一步，先排序。
//就可以用第二个数字减去第一个数字，如果差值大于1那么就说明这个中间就少了个数字。
//也就是我们要找的那个nums[i]-1。
//3,0,1
//0,1,3
func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			return nums[i] - 1
		}
	}
	return -1
}
