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

//数组
//值  下标
// 2   0
// 7   1
// 11  2
// 15  3

//map
//key value
// 2   0
// 7   1
// 11  2
// 15  3
//第一步，把数组里的值都放到map里去。以数组的值为key,数据的下标为value。
//第二步，遍历map，看是否能找到key值为target-value相对应的值。
//如果找到了，那么这个找到的值和此时遍历的value值就是我们需要的结果了。
func twoSum(nums []int, target int) []int {
	//hashTable := map[int]int{}
	hashTable := make(map[int]int)
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


//func twoSum(nums []int, target int) []int {
//	// 申请map
//	m := make(map[int]int)
//	// 遍历
//	for index, num := range nums {
//		// 到map中去查找目标2
//		key, ok := m[target - num]
//		// 找不到直接保存目标2
//		if !ok{
//			m[num] = index
//		}else {
//			// 找到直接返回目标1、目标2的下标
//			return []int{index, key}
//		}
//	}
//	return nil
//}
//