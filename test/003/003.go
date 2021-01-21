/*
  author='du'
  date='2020/11/6 4:31'
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

//abcabcbb
//a   //left不动，right++
//ab
//abc
// bca 移除窗口中的第一个元素，left++,right++

func main() {
	//var s="abcabcbb"
	//fmt.Println(lengthOfLongestSubstring(s))
	s1 := "abcbabc"
	fmt.Println(lengthOfLongestSubstring(s1))
}

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	window := s[left:right]
	res := 0
	for ; right < len(s); right++ {
		if index := strings.IndexByte(window, s[right]); index != -1 { //表示存在
			left += left + 1
		}
		window = s[left : right+1]
		//if len(window)>res{
		//	res=len(window)
		//}
		res = int(math.Max(float64(res), float64(len(window))))
	}
	return res
}
