/*
  author='du'
  date='2020/9/10 5:42'
*/
package main

import "fmt"

func main() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	fmt.Printf("before reverse s : %p , %v\n", &s, s)
	reverse(s)
	fmt.Println(s)
	fmt.Printf("after reverse s : %p , %v\n", &s, s)
}

func reverse(s []int) {
	fmt.Printf("s : %p , %v\n", &s, s)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
	fmt.Printf("s : %p , %v\n", &s, s)
}
