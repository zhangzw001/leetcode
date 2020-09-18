package main

import "fmt"

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
通过次数97,526提交次数160,225

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func permuteUnique(nums []int) [][]int {
	numsAll := make([][]int,len(nums)* len(nums) )
	if len(nums) ==0 {
		return numsAll
	}

	for i:=0; i < len(nums) ; i ++ {
		for j:=0; j < len(nums) ; j ++ {
			numsAll[i+j] = nums
			if i == j {
				continue
			}
			nums[i],nums[j] = nums[j],nums[i]
		}
	}
	return numsAll
}

func main() {
	fmt.Println(permuteUnique([]int{1,2,1}))
}