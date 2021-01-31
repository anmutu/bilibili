/*
  author='du'
  date='2021/1/24 10:00'
  github='https://github.com/anmutu'
*/
package main

import "fmt"

//链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
//
//给定一个字符串s和一个非空字符串p，找到s中所有是p的字母异位词的子串，返回这些子串的起始索引。
//字符串只包含小写英文字母，并且字符串s和 p的长度都不超过 20100。
//说明：
//字母异位词指字母相同，但排列不同的字符串。
//不考虑答案输出的顺序。
//
//示例1:
//输入:
//s: "cbaebabacd" p: "abc"
//输出:
//[0, 6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	var res []int
	for right < len(s) {
		window := s[left:right]
		if right-left < len(p) {
			right++
		} else if right-left == len(p) {
			if isAnagram(window, p) {
				res = append(res, left)
			}
			left++
			right++
		}
	}
	return res
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var c1, c2 [26]int
	for _, v := range s {
		c1[v-'a']++
	}
	for _, v := range t {
		c2[v-'a']++
	}
	return c1 == c2
}
