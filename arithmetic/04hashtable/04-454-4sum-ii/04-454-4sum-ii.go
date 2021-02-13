/*
  author='du'
  date='2021/2/10 1:37'
  github='https://github.com/anmutu'
*/

//链接：https://leetcode-cn.com/problems/4sum-ii
//给定四个包含整数的数组列表A , B , C , D ,
//计算有多少个元组 (i, j, k, l)，使得A[i] + B[j] + C[k] + D[l] = 0。
//为了使问题简单化，所有的 A, B, C, D 具有相同的长度N，且 0 ≤ N ≤ 500 。
//所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过231 - 1 。
//输入:
//A = [ 1, 2]
//B = [-2,-1]
//C = [-1, 2]
//D = [ 0, 2]
//输出:
//2
//解释:
//两个元组如下:
//1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
//2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0

package main

import "fmt"

func main() {
	n1, n2, n3, n4 := []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}
	fmt.Println(fourSumCount(n1, n2, n3, n4))
}

//暴力解法：500个。n(O^4)=>625,0000,0000
//用三个数组组成一个map查找表，再用另外一个数组去查找：n(O^3)=>1,2500,0000
//用两个数组组成一个map查找表，再用另外两个数组的组合去查找判断：n(O^2)=>25,0000
//key是组合的数的和，value就是相应的数量。
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
