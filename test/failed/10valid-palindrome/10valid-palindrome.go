/*
  author='du'
  date='2020/9/26 21:09'
*/
package main

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/valid-palindrome/
//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//说明：本题中，我们将空字符串定义为有效的回文串。
//输入: "A man, a plan, a canal: Panama"
//输出: true

//录屏失败。没有暂停了，而是结束了。

func main() {
	var s string = "A,b,a"
	fmt.Println(isPalindrome(s))
}

//用双指针的思路。
//会定义一个左右指针。
//从最左和最右开始比较，如果两个相等，那么就继续去比较，只到要不相等的。
//不相等了就返回false。全都相等就返回true。
//1.另外，要注意如果遇到了非字母或者非数字的符号要跳过。
//2.因为不考虑大小写，所以在比较前可以全部转为大写或者小写。
//条件：right>left
func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	s = strings.ToLower(s)
	for right > left {
		if !(s[left] > 'a' && s[left] < 'z' || s[left] > '0' && s[left] < '9') {
			left += 1
			continue
		}
		if !(s[right] > 'a' && s[right] < 'z' || s[right] > '0' && s[right] < '9') {
			right -= 1
			continue
		}
		if s[left] == s[right] {
			left += 1
			right -= 1
		} else {
			return false
		}
	}
	return true
}
