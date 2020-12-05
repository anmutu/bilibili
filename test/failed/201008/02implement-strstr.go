/*
  author='du'
  date='2020/10/8 4:53'
*/
package main

import "fmt"

func main() {
	long, short := "hello", "ll"
	fmt.Println(strStr(long, short))
}

//明确遍历：遍历的是长串。
//如何找出长串里与短串相对应的串：下标。
//遍历条件：len(long)-len(short)。
//特殊情况的处理。
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
			return 0
		} else {
			p += 1
		}
	}
	return -1
}
