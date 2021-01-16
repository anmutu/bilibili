/*
  author='du'
  date='2021/1/12 0:44'
*/
package main

import "fmt"

func main() {
	deferCall()
	fmt.Println("main 正常结束")
}

func deferCall() {
	defer func() { fmt.Println("defer: panic 之前1") }()
	defer func() { fmt.Println("defer: panic 之前2") }()
	panic("异常内容") //触发defer出栈
	defer func() { fmt.Println("defer: panic 之后，永远执行不到") }()
}

//此处用2。
//遇到panic时，遍历本协程的defer链表，并执行defer。
//1.在执行defer过程中:遇到recover则停止panic，返回recover处继续往下执行。
//2.如果没有遇到recover，遍历完本协程的defer链表后，則抛出panic信息。
