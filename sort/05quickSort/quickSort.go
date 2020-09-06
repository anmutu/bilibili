/*
  author='du'
  date='2020/9/1 0:47'
*/
package main

import "fmt"

func main() {
	array := []int{1, 8, 0, 10, 2, 7, 6}
	quickSort(array)
}

func quickSort(array []int) {
	if len(array) <= 1 {
		fmt.Println(array)
	}
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(array); i++ {
		if array[i] < array[0] {
			left = append(left, array[i])
		} else {
			right = append(right, array[i])
		}
	}
	fmt.Println(left)
	fmt.Println(array[0])
	fmt.Println(right)
	fmt.Println(array)
}
