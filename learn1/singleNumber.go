package main

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions/xm0u83/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

func singleNumber2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	status := make(map[int]bool)
	for _, v := range nums {
		if status[v] {
			status[v] = false
		}else {
			status[v] = true
		}
	}

	for _, v := range nums {
		if status[v] {
			return v
		}
	}
	return 0
}

// 所有的元素异或运算之后 等于 唯一的那个数
// (a1 ^ a1 ) ^ ( a2 ^ a2 ) ^ ... (an ^ an) ^ a(n+1) = a(n+1)
func singleNumber(nums []int) int {
	status := 0
	for _, v := range nums {
		status ^= v
	}
	return status
}
func main() {
	fmt.Println(singleNumber([]int{-1,2,2}))
	fmt.Println(singleNumber([]int{2,2,11,3,3,42,42}))
}