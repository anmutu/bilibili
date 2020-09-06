/*
  author='du'
  date='2020/9/3 21:17'
*/

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
//给定一个数组，它的第i个元素是一支给定股票第i天的价格。
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

//example
//输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

package main

import "fmt"

func main() {
	var nums []int = []int{7, 1, 5, 3, 6, 4}
	res := bestTime2BuyOrSell(nums)
	fmt.Println(res)
}

func bestTime2BuyOrSell(nums []int) int {
	maxValue := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			maxValue += nums[i+1] - nums[i]
		}
	}
	return maxValue
}
