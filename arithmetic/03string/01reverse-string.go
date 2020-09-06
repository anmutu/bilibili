/*
  author='du'
  date='2020/9/4 13:23'
*/

//# https://leetcode-cn.com/problems/reverse-string/
//
//# 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
//# 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
//# 你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
//
//# example
//# 输入：["h","e","l","l","o"]
//# # 输出：["o","l","l","e","h"]

package main

import "fmt"

func main() {
	var s []byte = []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))
}

//双指针解决。
func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
