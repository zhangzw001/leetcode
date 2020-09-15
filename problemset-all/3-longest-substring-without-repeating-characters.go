package main

import "fmt"

//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	var sum int
	left , right := 0,0
	var status []int
	for left < len(s)-1 {


		if status[s[right]] != 0 {
			right++

		}else {
			right --
			fmt.Println(left,right)
			if right - left > sum {
				sum = right - left
			}
			left++
			right = left+1
		}


	}
	return sum
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		fmt.Println(bitSet)
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将X标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}
func main() {
	a := "pwwkew"
	fmt.Println(lengthOfLongestSubstring2(a))
}
// [p,w],[p,w,w] ->退出
// [w,w] -> 退出
// [w,k] ,[w,k,e], [w,k,e,w] ->退出