/*
  author='du'
  date='2021/1/18 23:28'
  github='https://github.com/anmutu'
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

//给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
//
//示例1:
//url：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	window := s[left:right]
	res := 0
	for ; right < len(s); right++ {
		if index := strings.IndexByte(window, s[right]); index != -1 { //表示存在
			left += index + 1
		}
		window = s[left : right+1]
		res = int(math.Max(float64(res), float64(len(window))))
	}
	return res
}
