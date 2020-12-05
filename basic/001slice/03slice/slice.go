/*
  author='du'
  date='2020/9/10 5:51'
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
	s = append(s, 999, 1000, 1001)
	//fmt.Println("after append:",s,"len():",len(s),"cap():",cap(s))
	//fmt.Printf("s : %p , %v\n", &s, s)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
		//fmt.Println(i,j,s[i],s[j])
	}
	//fmt.Println("s",s)
}

//在此示例中，我们添加了三个元素，而我们的切片没有足够的容量。
//而是分配了一个新数组，我们更新的切片指向该数组。当我们最终开始反转切片中的元素时，它不再影响我们的初始数组，
//而是在完全不同的数组上运行。
