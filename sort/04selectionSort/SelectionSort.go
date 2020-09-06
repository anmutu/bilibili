/*
  author='du'
  date='2020/9/1 0:31'
*/
package main

import "fmt"

func main() {
	array := []int{1, 8, 0, 10, 2, 7, 6}
	selectionSort(array)
}

func selectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println(array)
}
