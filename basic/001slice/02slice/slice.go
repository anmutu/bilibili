/*
  author='du'
  date='2020/9/10 5:48'
*/
package main

import "fmt"

func main() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	//fmt.Println("before reverse",s,"len():",len(s),"cap():",cap(s))
	//fmt.Printf("s : %p , %v\n", &s, s)
	reverse(s)
	//fmt.Println("after reverse",s,"len():",len(s),"cap():",cap(s))
	fmt.Printf("s : %p , %v\n", &s, s)
}

func reverse(s []int) {
	//fmt.Println("before append:",s,"len():",len(s),"cap():",cap(s))
	//fmt.Printf("s : %p , %v\n", &s, s)
	s = append(s, 999)
	//fmt.Println("after append:",s,"len():",len(s),"cap():",cap(s))
	//fmt.Printf("s : %p , %v\n", &s, s)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//当我们调用时，append将创建一个新的切片。
//新切片具有新length属性，该属性不是指针，但仍指向同一数组。
//因此，我们其余的代码最终会反转原始切片所引用的数组，但是原始切片的长度没有改变。
