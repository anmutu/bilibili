/*
  author='du'
  date='2021/1/12 0:37'
*/
package main

import "fmt"

func main() {
	returnAndDefer()
}

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

//结论为：return之后的语句先执行，defer后的语句后执行
