/*
  author='du'
  date='2021/1/11 23:55'
*/
package main

import "fmt"

//不难看出，Foo()的形参x interface{}是一个空接口类型eface struct{}。
//x 结构体本身不为nil，而是data指针指向的p为nil。

func main() {
	var p *int = nil
	Foo(p)
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
