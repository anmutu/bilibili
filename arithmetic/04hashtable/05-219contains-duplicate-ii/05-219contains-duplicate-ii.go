/*
  author='du'
  date='2021/2/13 20:14'
  github='https://github.com/anmutu'
*/
package main

import "fmt"

func main() {
	//nums:=[]int{1,2,3,1}
	//k:=3
	//fmt.Println(containsNearbyDuplicate(nums,k))
	//n1:=[]int{1,2,3,1,2,3}
	//k1:=2
	//fmt.Println(containsNearbyDuplicate(n1,k1))
	n2 := []int{4, 6, 1, 2, 3, 1, 5}
	k2 := 3
	fmt.Println(containsNearbyDuplicate(n2, k2))
}

//定义查找表：以数组元素的值为key，存在则设置其value为true。
//在遍历的同时，使查找表最多保持k的长度。
//一旦超过k，则将最面前的元素删除掉。
func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]bool)
	for i, v := range nums {
		fmt.Println(hash, k)
		if _, ok := hash[v]; ok {
			return true
		}
		hash[v] = true
		if len(hash) > k {
			fmt.Println("delete", nums[i-k])
			delete(hash, nums[i-k])
		}
	}
	return false
}
