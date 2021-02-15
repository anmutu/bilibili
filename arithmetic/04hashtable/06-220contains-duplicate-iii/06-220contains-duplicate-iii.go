/*
  author='du'
  date='2021/2/13 20:14'
  github='https://github.com/anmutu'
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 1}
	k, t := 3, 0
	fmt.Println(containsNearbyDuplicate(nums, k, t))
}

func containsNearbyDuplicate(nums []int, k int, t int) bool {
	for i := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) <= float64(t) {
				return true
			}
		}
	}
	return false
}
