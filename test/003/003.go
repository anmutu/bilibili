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
//a   //left不动，right++   =>ab
//ab  //left不动，right++   =>abc
//abc                      =>abca
//abca //移除窗口中的第一个元素，left++,right++     =>bca  =>纪录3
// bca
func main() {
	//var s="abcabcbb"
	//fmt.Println(lengthOfLongestSubstring(s))
	s1 := "abcbabc"
	fmt.Println(lengthOfLongestSubstring(s1))
}

/*
 定义一个左右指针，left和right
 使可得到窗口window[left:right]
 判断windows中是否有s[right]元素
 如果存在:
 则需要将left的位置变更为当前下标位置+1，
 且right也+1,
 取此时的window，判断其大小是不是每次最大的。
 如果不存在：
 则需要将right往后移一步。
*/

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
