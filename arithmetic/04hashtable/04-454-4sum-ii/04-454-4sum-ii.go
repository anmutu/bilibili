/*
  author='du'
  date='2021/2/10 1:37'
  github='https://github.com/anmutu'
*/
package main

import "fmt"

func main() {
	n1, n2, n3, n4 := []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}
	fmt.Println(fourSumCount(n1, n2, n3, n4))
}

// 暴力解法：n(O^4)=>625,0000,0000=>pass
//n(O^3)=>1,2500,0000=>pass
func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	hash := make(map[int]int)
	for _, v1 := range A {
		for _, v2 := range B {
			hash[v1+v2]++
		}
	}
	fmt.Println(hash)
	for _, v3 := range C {
		for _, v4 := range D {
			if _, ok := hash[0-v3-v4]; ok {
				res += 1
			}
		}
	}
	return res
}
