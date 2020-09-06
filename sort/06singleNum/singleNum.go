/*
  author='du'
  date='2020/9/1 3:06'
*/
package main

import (
	"fmt"
	"sort"
)

//一个数组，除某个元素只出现1次外，其余均出现2次。找出这个只出现1次的元素。

func main() {
	nums := []int{2, 2, 1}
	res := singleNumber(nums)
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	i := 0
	for i < len(nums)-2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		} else {
			i += 2
		}
	}
	return nums[len(nums)]
}
