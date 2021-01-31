/*
  author='du'
  date='2021/1/24 10:03'
  github='https://github.com/anmutu'
*/
package main

import "fmt"

//链接：https://leetcode-cn.com/problems/valid-anagram
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//示例1:
//输入: s = "anagram", t = "nagaram"
//输出: true

/*
用map纪录单词出现的次数，然后比较map是否相等。
如何去存次单词出现的次数？
既然都是小写。
'a'的ASCII的值是97。
a到z是97到122。
既然全是小写，那么就可以设置一个长度为26的map了。
所以遍历字符串让其中的value-97就可以得到相应单词了。
匹配则把这个值+1,以计算单词出现的个数。
*/
func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
	s1 := "cat"
	s2 := "car"
	fmt.Println(isAnagram(s1, s2))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//m1,m2:=[26]int{},[26]int{}
	var m1, m2 [26]int
	for _, v := range s {
		m1[v-'a']++
	}
	for _, v := range t {
		m2[v-'a']++
	}
	return m1 == m2
}
