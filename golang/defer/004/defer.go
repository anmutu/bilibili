/*
  author='du'
  date='2021/1/12 0:41'
*/
package main

import "fmt"

func main() {
	fmt.Println(returnButDefer())
}

func returnButDefer() (t int) { //由003可知t初始化0， 并且作用域为该函数全域
	defer func() {
		t = t * 10
	}()
	return 1
}

//该returnButDefer()本应的返回值是1，但是在return之后，
//又被defer的匿名func函数执行，所以t=t*10被执行，最后returnButDefer()返回给上层main()的结果为10
