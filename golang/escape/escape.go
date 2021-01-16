/*
  author='du'
  date='2021/1/9 1:30'
*/
package main

import "fmt"

func main() {
	mainVal := foo1(888)
	fmt.Println(*mainVal, mainVal)
}

func foo1(argVal int) *int {
	var fooVal1 int = 11
	var fooVal2 int = 12
	var fooVal3 int = 13
	var fooVal4 int = 14
	var fooVal5 int = 15

	//此处循环是防止go编译器将foo优化成inline(内联函数)
	//如果是内联函数，main调用foo将是原地展开，所以fooVal1-5相当于main作用域的变量
	//即使foo_val3发生逃逸，地址与其他也是连续的
	for i := 0; i < 5; i++ {
		println(&argVal, &fooVal1, &fooVal2, &fooVal3, &fooVal4, &fooVal5)
	}

	//返回fooVal3给main函数
	return &fooVal3
}
