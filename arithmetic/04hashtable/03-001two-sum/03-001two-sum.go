/*
  author='du'
  date='2021/2/10 1:35'
  github='https://github.com/anmutu'
*/
package main

import "fmt"

func main() {
	var nums []int = []int{2, 11, 15, 7}
	target := 9
	fmt.Println(twoSum(nums, target))

	//会覆盖情况的举例。
	var n1 = []int{4, 2, 4, 7}
	t1 := 8
	fmt.Println(twoSum1(n1, t1))
	fmt.Println(twoSum(n1, t1))
}

//查找map表设计：以数组value为key，index为值。
//判断key为target-value的值是否存在于map中。
//如果存在则将此时的索引和这个值返回。
//注意：需要查找表需要动态往后增大。
func twoSum(nums []int, target int) []int {
	var res []int
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
