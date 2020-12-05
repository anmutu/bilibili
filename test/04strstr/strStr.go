/*
  author='du'
  date='2020/10/8 3:29'
*/
package main

import "fmt"

//https://leetcode-cn.com/problems/implement-strstr/
//给定一个haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
//输入: haystack = "hello", needle = "ll"
//输出: 2

func main() {
	short, long := "ll", "hello"
	fmt.Println(strStr(long, short))
}

//遍历长的。
//怎么取与短串对应长度的值。
//遍历的条件是什么。
func strStr(long string, short string) int {
	if len(short) > len(long) {
		return -1
	}
	if short == long {
		return 0
	}
	p := 0
	length := len(short)
	for p <= len(long)-len(short) {
		if long[p:p+length] == short {
			return p
		} else {
			p += 1
		}
	}
	return -1
}
