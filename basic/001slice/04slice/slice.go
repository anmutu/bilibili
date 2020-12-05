/*
  author='du'
  date='2020/9/10 6:44'
*/
package main

import "fmt"

func main() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	reverse(s)
	fmt.Println(s)
}

func reverse(s []int) {
	newElem := 999
	for len(s) < cap(s) {
		fmt.Println("Adding an element:", newElem, "cap:", cap(s), "len:", len(s))
		s = append(s, newElem)
		newElem++
	}
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
