/*
  author='du'
  date='2021/2/9 21:18'
  github='https://github.com/anmutu'
*/

package main

import "fmt"

func main() {
	var n1 []int = []int{1, 2, 2, 1}
	var n2 []int = []int{2, 2}
	fmt.Println(interSectionOf2ArraysMap(n1, n2))
}

func interSectionOf2ArraysMap(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	hash := make(map[int]int)
	for _, v := range nums1 {
		hash[v] += 1
	}
	for _, v := range nums2 {
		if hash[v] > 0 {
			res = append(res, v)
			hash[v]--
		}
	}
	return res
}
