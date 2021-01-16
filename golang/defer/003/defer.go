/*
  author='du'
  date='2021/1/12 0:38'
*/
package main

import "fmt"

func main() {
	DeferFunc(10)
}

func DeferFunc(i int) (t int) {
	fmt.Println("t = ", t)
	return 2
}

//证明，只要声明函数的返回值变量名称，就会在函数初始化时候为之赋值为0，而且在函数体作用域可见。
