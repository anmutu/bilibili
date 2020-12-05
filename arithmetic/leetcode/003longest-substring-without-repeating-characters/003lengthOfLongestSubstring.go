/*
  author='du'
  date='2020/11/6 5:20'
*/
package main

import (
	"fmt"
)

//给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

//原题链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//代码链接：https://github.com/anmutu/bilibili

func main() {
	var s="abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

//abcabcbb
//a   1
//ab  2  left不变，right++
//abc 3
// bca 3  left++,right++
func lengthOfLongestSubstring(s string) int {
	//left,right:=0,0
	//res:=0
	//for right<len(s){
  //
	//}
  //
	//for ;
  //return res
	freq := make([]int,128)
	var res = 0
	start,end := 0,-1
	for start<len(s){
		fmt.Println(s[end+1])
		if end+1<len(s)&&freq[s[end+1]] == 0{
			end++
			freq[s[end]]++
		}else{
			freq[s[start]]--
			start++
		}
		res = max(res,end-start+1)
	}
	return res


}


func max(i,j int)int{
	if i>j{
		return i
	}else{
		return j
	}
}