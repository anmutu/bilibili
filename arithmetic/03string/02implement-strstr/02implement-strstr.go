/*
  author='du'
  date='2020/10/8 5:05'
*/
package main

import "fmt"

func main() {
	long, short := "hello", "ll"
	fmt.Println(strStr(long, short))
}

//1.明确遍历的是长串。
//2.如何在长串里找到与短串长度相对应的值。下标。
//3.条件：len(long)-len(short)
//4.特殊情况的处理。
func strStr(long string, short string) int {
	if len(long) < len(short) {
		return -1
	}
	if long == short {
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
