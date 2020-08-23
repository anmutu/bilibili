/*
  author='du'
  date='2020/8/16 21:44'
*/
package main

import (
	"fmt"
	"math"
)

var pyramid [4][]int=[4][]int{
	{1},  //第一层
	{2,3},  //第二层,也会有2个数。下标也会比层数少1。
	{4,5,6}, //第三层,也会有3个数。下标会比层数少1。
	{7,8,9,10}, //第四层
}


func main() {
	//fmt.Println(Test(pyramid))
	fmt.Println(getBest(pyramid))
}

func getBest([4][]int) int{
	//外层的这个循环是需要遍历n-1次的，n是这个金字塔的一个层数。
	for j:=3;j>0;j--{
		for i:=0;i<j;i++{
			fmt.Println("before",pyramid[2][i])
			fmt.Println("lower left",pyramid[3][i])
			fmt.Println("lower right",pyramid[3][i+1])
			pyramid[j-1][i] += int(math.Max(float64(pyramid[j][i]),float64(pyramid[j][i+1])))
			fmt.Println("after",pyramid[2][i])
			fmt.Println("+++++++++++++++++++++++++")
		}
	}
	return pyramid[0][0]
}

func Test([4][]int) int{
	//把第三层对第四层的这个最优值的问题解决。
	for i:=0;i<3;i++{
		fmt.Println("before",pyramid[2][i])
		fmt.Println("lower left",pyramid[3][i])
		fmt.Println("lower right",pyramid[3][i+1])
		pyramid[2][i] += int(math.Max(float64(pyramid[3][i]),float64(pyramid[3][i+1])))
		fmt.Println("after",pyramid[2][i])
		fmt.Println("+++++++++++++++++++++++++")
	}
	fmt.Println(pyramid[2])

	fmt.Println("**************************************************")
	//把第二层最优化
	for i:=0;i<2;i++{
		fmt.Println("before",pyramid[2-1][i])
		fmt.Println("lower left",pyramid[3-1][i])
		fmt.Println("lower right",pyramid[3-1][i+1])
		pyramid[2-1][i] += int(math.Max(float64(pyramid[3-1][i]),float64(pyramid[3-1][i+1])))
		fmt.Println("after",pyramid[2-1][i])
		fmt.Println("+++++++++++++++++++++++++")
	}
	fmt.Println(pyramid[1])


	fmt.Println("**************************************************")
	//把第一层最优化
	for i:=0;i<1;i++{
		fmt.Println("before",pyramid[2-1-1][i])
		fmt.Println("lower left",pyramid[3-1-1][i])
		fmt.Println("lower right",pyramid[3-1-1][i+1])
		pyramid[2-1-1][i] += int(math.Max(float64(pyramid[3-1-1][i]),float64(pyramid[3-1-1][i+1])))
		fmt.Println("after",pyramid[2-1-1][i])
		fmt.Println("+++++++++++++++++++++++++")
	}
	fmt.Println(pyramid[0])
	return pyramid[0][0]
}