package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	lenS := len(s)
	lenP := len(p)
	if lenS < lenP {
		return []int{}
	}
	need := [26]int{}
	window := [26]int{}
	for i := 0; i < lenP; i++ {
		need[p[i]-'a']++
	}
	check := func() bool {
		for i := 0; i < 26; i++ {
			if window[i] != need[i] {
				return false
			}
		}
		return true
	}
	match := 0          // 表示匹配的字符个数
	left, right := 0, 0 // 左开右避
	res := []int{}
	var idx, leftidx byte
	for right < lenS {
		idx = s[right] - 'a'
		window[idx]++
		right++
		if window[idx] <= need[idx] {
			match++
		}
		for match == lenP {
			if check() {
				res = append(res, left)
			}
			leftidx = s[left] - 'a'
			left++
			// 这里要注意一下细节,如果need[leftChar]==0，也可能导致match被减
			if window[leftidx] <= need[leftidx] {
				match--
			}
			window[leftidx]--
		}
	}
	return res
}
