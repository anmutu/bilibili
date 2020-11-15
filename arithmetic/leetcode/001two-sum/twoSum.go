/*
  author='du'
  date='2020/11/14 12:31'
*/
package main

import "fmt"

func main() {
	var nums []int = []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

//值  下标
//key value
// 2   0
// 7   1
// 11  2
// 15  3
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for index, value := range nums {
		//if p,ok:=hashTable[target-value];ok{
		//	return []int{p,index}
		//}
		hashTable[value] = index
	}

	for k, v := range hashTable {
		if p, ok := hashTable[target-k]; ok {
			return []int{p, v}
		}
	}

	return nil
}
