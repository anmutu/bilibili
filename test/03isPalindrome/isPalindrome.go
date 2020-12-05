/*
  author='du'
  date='2020/9/26 16:33'
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	//var s string="A man, a plan, a canal: Panama"
	//fmt.Println(isPalindrome(s))

	var s1 string = "ab,a"
	fmt.Println(isPalindrome(s1))

}

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
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
