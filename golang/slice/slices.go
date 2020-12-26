/*
  author='du'
  date='2020/12/26 8:46'
*/
package main

import "fmt"

func main() {
	getSlicesValue()
	slicesIsViewOfArray()
	extendingSlices()
	append2Slices()
	createSlices()
}

//切片取值
func getSlicesValue() {
	fmt.Println("1.进入getSlicesValue()函数")
	arr := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])
}

//证明下slices里array的一个view,也就是修改slices里面的值也就是修改array里面的值了。
//slices本身是没有数据的，是对底层array的一个view
func slicesIsViewOfArray() {
	fmt.Println("2.进入getSlicesValue()函数")
	arr := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:]
	fmt.Println("update前的s1,即arr[2:]:", s1)
	updateSlices(s1)
	fmt.Println("update过后的s1:", s1) //把s1第一个元素改成了1000
	fmt.Println("现在arr的值是:", arr)   //发现最原始的2号位变成了1000
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
	arr := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("arr最初的值是:%v,len(arr)=%d,cap(arr)=%d \n", arr, len(arr), cap(arr))
	fmt.Printf("s1的值，即arr[2:6]的值是:%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2的值，即s2[3:5]的值是:%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))
}

//进入append2Slices()函数
//s1,s2,s3,s4= [6 7 8] [6 7 8 9] [6 7 8 9 10] [6 7 8 9 10 11]
//arr= [0 1 2 3 4 5 6 7 8]
//超过cap会分配更大的底层数组，已经不是同一个了。
func append2Slices() {
	fmt.Println("4.进入append2Slices()函数")
	arr := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
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

//cap的长度的规律。当len为5时，则cap则为4的2倍。以此类推。
//当len超过原来len的长度时一直是2*len吗？
func createSlices() {
	fmt.Println("5.进入createSlices()函数")
	var s []int
	for i := 0; i < 20; i++ {
		printSlices(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
}

func printSlices(s []int) {
	fmt.Printf("%v,len(s)=%d,cap(s)=%d \n", s, len(s), cap(s))
}

func updateSlices(s []int) {
	s[0] = 1000
}
