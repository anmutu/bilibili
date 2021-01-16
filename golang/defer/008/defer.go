/*
  author='du'
  date='2021/1/12 1:31'
  github='https://github.com/anmutu'
*/
package main

import "fmt"

func main() {
	defer function(1, function(3, 0))
	defer function(2, function(4, 0))
}

func function(index int, value int) int {
	fmt.Println(index)
	return index
}

//这里，有4个函数，他们的index序号分别为1，2，3，4。
//那么这4个函数的先后执行顺序是什么呢？这里面有两个defer，
//所以defer一共会压栈两次，先进栈1，后进栈2。 那么在压栈function1的时候，需要连同函数地址、函数形参一同进栈，
//那么为了得到function1的第二个参数的结果，所以就需要先执行function3将第二个参数算出，
//那么function3就被第一个执行。同理压栈function2，就需要执行function4算出function2第二个参数的值。
//然后函数结束，先出栈fuction2、再出栈function1.
//
//所以顺序如下：
//defer压栈function1，压栈函数地址、形参1、形参2(调用function3) --> 打印3
//defer压栈function2，压栈函数地址、形参1、形参2(调用function4) --> 打印4
//defer出栈function2, 调用function2 --> 打印2
//defer出栈function1, 调用function1--> 打印1
