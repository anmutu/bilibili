/*
  author='du'
  date='2021/2/10 1:35'
  github='https://github.com/anmutu'
*/

//# 给定一个整数数组 nums和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
//# 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//# example
//# 给定 nums = [2, 7, 11, 15], target = 9
//# 因为 nums[0] + nums[1] = 2 + 7 = 9
//# 所以返回 [0, 1]

package main

import "fmt"

func main() {
	var nums []int = []int{2, 11, 15, 7}
	target := 9
	fmt.Println(twoSum(nums, target))
	fmt.Println("+++++++++++++++++++")
	var n1 = []int{4, 2, 4, 7}
	t1 := 8
	fmt.Println(twoSum1(n1, t1))
	fmt.Println("+++++++++++++++++++")
	fmt.Println(twoSum(n1, t1))
}

//设计查找表：以数组的value为key,index为值。
//在查找表里去找：以target-value的值是否存在于map中。
//如果在的话那么就找到了，这个索引和相应的值返回就好。
//注意：查找表是需要动态的增大的。
func twoSum(nums []int, target int) (res []int) {
	hash := make(map[int]int)
	for index, value := range nums {
		if p, ok := hash[target-value]; ok {
			res = []int{p, index}
		}
		hash[value] = index
		fmt.Println(value, hash)
	}
	return res
}

//把所有的元素都放入查找表中会导致相同的元素覆盖的问题。
//所以这个方法是不可取的。需要注意的。
func twoSum1(nums []int, target int) (res []int) {
	hash := make(map[int]int)
	for index, value := range nums {
		hash[value] = index
	}
	fmt.Println(hash)

	for index, value := range nums {
		if p, ok := hash[target-value]; ok {
			res = []int{p, index}
		}
	}
	return res
}
