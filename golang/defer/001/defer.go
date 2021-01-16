/*
  author='du'
  date='2021/1/12 0:36'
*/
package main

import "fmt"

func main() {
	defer func1()
	defer func2()
	defer func3()
}

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}
