/*
  author='du'
  date='2020/9/1 0:12'
*/
package main

import "fmt"

func main() {
	array := []int{1, 8, 0, 10, 2, 7, 6}
	BobbleSort(array)
}

func BobbleSort(ints []int) {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i] > ints[j] {
				ints[i], ints[j] = ints[j], ints[i]
			}
		}
	}
	fmt.Println(ints)
}
