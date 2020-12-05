/*
  author='du'
  date='2020/9/8 13:58'
*/

package main

import (
	"fmt"
)

func main() {
	getSlicesValue()
	fmt.Println()
	slicesIsViewOfArray()
	fmt.Println()
	extendingSlices()
	fmt.Println()
	createSlices()
	fmt.Println()
}

//切片取值
func getSlicesValue() {
	fmt.Println("1.进入getSlicesValue()函数")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])
}

//证明下slices里array的一个view,也就是修改slices里面的值也就是修改array里面的值了。
//slices本身是没有数据的，是对底层array的一个view
func slicesIsViewOfArray() {
	fmt.Println("2.进入getSlicesValue()函数")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:]
	fmt.Println("update前的s1,即arr[2:]:", s1)
	updateSlices(s1)
	fmt.Println("update过后的s1:", s1) //把s1第一个元素(也就是2这个元素)改成了1000
	fmt.Println("现在arr的值是:", arr)
}

//进入extendingSlices()函数
//arr最初的值是: [0 1 2 3 4 5 6 7 8]
//s1的值，即arr[2:6]的值是: [2 3 4 5]
//s2的值，即s1[3:5]的值是: [5 6]
//为什么s2的值是会有，因为slices是可以扩展的。
// 一个slices里面有三个变量：
// 1.ptr。指向第一个，
// 2.len。表示这个slices的长度，
// 3.cap。指从ptr到array的最后一个的长度，只要不超过这个cap的长度，就都可以扩展。
// 4.slice是可以向后扩展的，但是不能向前扩展。
func extendingSlices() {
	fmt.Println("3.进入extendingSlices()函数")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("arr最初的值是:%v,len(arr)=%d,cap(arr)=%d \n", arr, len(arr), cap(arr))
	fmt.Printf("s1的值，即arr[2:6]的值是:%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2的值，即s1[3:5]的值是:%v,len(s1)=%d,cap(s1)=%d\n", s2, len(s2), cap(s2))
}

//cap的长度的规律。当len为5时，则cap则为4的2倍。以此类推。
//比如长度是1035呢？那么cap会是多少呢？
func createSlices() {
	fmt.Println("4.进入createSlices()函数")
	var s []int
	for i := 0; i < 1035; i++ {
		printSlices(s)
		s = append(s, 2*i+1)
	}
}

func printSlices(s []int) {
	fmt.Printf("%v,len(s)=%d,cap(s)=%d \n", s, len(s), cap(s))
}

func updateSlices(s []int) {
	s[0] = 1000
}

/*

1.进入getSlicesValue()函数
arr[2:6]= [2 3 4 5]
arr[:6]= [0 1 2 3 4 5]
arr[2:]= [2 3 4 5 6 7 8]
arr[:]= [0 1 2 3 4 5 6 7 8]

2.进入getSlicesValue()函数
update前的s1,即arr[2:]: [2 3 4 5 6 7 8]
update过后的s1: [1000 3 4 5 6 7 8]
现在arr的值是: [0 1 1000 3 4 5 6 7 8]

3.进入extendingSlices()函数
arr最初的值是:[0 1 2 3 4 5 6 7 8],len(arr)=9,cap(arr)=9
s1的值，即arr[2:6]的值是:[2 3 4 5],len(s1)=4,cap(s1)=7
s2的值，即s1[3:5]的值是:[5 6],len(s1)=2,cap(s1)=4

4进入append2Slices()函数
[6 7 8],len(s)=3,cap(s)=3
[6 7 8 9],len(s)=4,cap(s)=6
[6 7 8 9 10],len(s)=5,cap(s)=6
[6 7 8 9 10 11],len(s)=6,cap(s)=6
[6 7 8 9 10 11 12],len(s)=7,cap(s)=12
arr= [0 1 2 3 4 5 6 7 8]

5.进入createSlices()函数
[],len(s)=0,cap(s)=0
[1],len(s)=1,cap(s)=1
[1 3],len(s)=2,cap(s)=2
[1 3 5],len(s)=3,cap(s)=4
[1 3 5 7],len(s)=4,cap(s)=4
[1 3 5 7 9],len(s)=5,cap(s)=8
[1 3 5 7 9 11],len(s)=6,cap(s)=8
[1 3 5 7 9 11 13],len(s)=7,cap(s)=8
[1 3 5 7 9 11 13 15],len(s)=8,cap(s)=8
[1 3 5 7 9 11 13 15 17],len(s)=9,cap(s)=16
[1 3 5 7 9 11 13 15 17 19],len(s)=10,cap(s)=16
[1 3 5 7 9 11 13 15 17 19 21],len(s)=11,cap(s)=16
[1 3 5 7 9 11 13 15 17 19 21 23],len(s)=12,cap(s)=16
[1 3 5 7 9 11 13 15 17 19 21 23 25],len(s)=13,cap(s)=16
[1 3 5 7 9 11 13 15 17 19 21 23 25 27],len(s)=14,cap(s)=16
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29],len(s)=15,cap(s)=16
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31],len(s)=16,cap(s)=16
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33],len(s)=17,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35],len(s)=18,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37],len(s)=19,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39],len(s)=20,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41],len(s)=21,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43],len(s)=22,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45],len(s)=23,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47],len(s)=24,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49],len(s)=25,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51],len(s)=26,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53],len(s)=27,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55],len(s)=28,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57],len(s)=29,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57 59],len(s)=30,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57 59 61],len(s)=31,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57 59 61 63],len(s)=32,cap(s)=32
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57 59 61 63 65],len(s)=33,cap(s)=64
[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57 59 61 63 65 67],len(s)=34,cap(s)=64

6.进入copyDelPop()函数
[0 0 0 0 0 0 0 0 0 0],len(s)=10,cap(s)=32
[0 1 2 3],len(s)=4,cap(s)=4
将s2copy到s1的结果为：
[0 1 2 3 0 0 0 0 0 0],len(s)=10,cap(s)=32
将下标为3的delete掉的结果为：
[0 1 2 0 0 0 0 0 0 0],len(s)=10,cap(s)=32
从头pop一个数据出去：
pop出去的头的数值为:0,pop后的结果为:[1 2 0 0 0 0 0 0 0][1 2 0 0 0 0 0 0 0],len(s)=9,cap(s)=31
从尾部pop一个数据出去：
pop出去的尾部的数值为:0,pop后的结果为:[1 2 0 0 0 0 0 0][1 2 0 0 0 0 0 0],len(s)=8,cap(s)=31

*/

//进入append2Slices()函数
//arr= [0 1 2 3 4 5 6 7 8]
//超过cap会分配更大的底层数组，已经不是同一个了。
func append2Slices() {
	fmt.Println("进入append2Slices()函数")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[6:]
	s2 := append(s1, 9)
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	printSlices(s1)
	printSlices(s2)
	printSlices(s3)
	printSlices(s4)
	printSlices(s5)
	fmt.Println("arr=", arr)
}

//删除就是用append两个slices,这两个slience不包括要删除的元素就行。
func copyDelPop() {
	fmt.Println("6.进入copyDelPop()函数")
	s1 := make([]int, 10, 32)
	s2 := []int{0, 1, 2, 3}
	printSlices(s1)
	printSlices(s2)
	fmt.Println("将s2copy到s1的结果为：")
	copy(s1, s2)
	printSlices(s1)
	//把现在下标为3的数值delete掉
	fmt.Println("将下标为3的delete掉的结果为：")
	s2 = append(s1[:3], s1[4:]...)
	printSlices(s1)
	fmt.Println("从头pop一个数据出去：")
	front := s1[0]
	s1 = s1[1:]
	fmt.Printf("pop出去的头的数值为:%v,pop后的结果为:%v", front, s1)
	printSlices(s1)
	fmt.Println("从尾部pop一个数据出去：")
	tail := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]
	fmt.Printf("pop出去的尾部的数值为:%v,pop后的结果为:%v", tail, s1)
	printSlices(s1)
}
