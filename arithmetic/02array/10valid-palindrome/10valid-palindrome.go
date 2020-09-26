/*
  author='du'
  date='2020/9/26 21:32'
*/
package main

import (
	"fmt"
	"strings"
)

// https://github.com/anmutu/bilibili
// https://leetcode-cn.com/problems/valid-palindrome/
//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//说明：本题中，我们将空字符串定义为有效的回文串。
//输入: "A man, a plan, a canal: Panama"
//输出: true

func main() {
	//aba
	var s string = "a,bA"
	fmt.Println(isPalindrome(s))
	var s1 string = "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s1))
}

//会定义一个左指针和一个右指针。
//一直去遍历这个string，如果左边的和右边的相等，就一直遍历。
//如果不相等了，那么就说明这个不是回文串了，就返回false。
//要注意1：需要先把这个字符串统一变成大写或者小写。
//要注意点2：遇到了非数字或者非字母的情况。
//2a.如果是左侧的遇到了，那么指针就往右移一步，继续比较。
//2b.如果是右侧遇到了，那么指针就往左移一步，继续比较。
//条件：right>left
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for right > left {
		if !(s[left] >= 'a' && s[left] <= 'z' || s[left] >= '0' && s[left] <= '9') {
			left += 1
			continue
		}
		if !(s[right] >= 'a' && s[right] <= 'z' || s[right] >= '0' && s[right] <= '9') {
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
