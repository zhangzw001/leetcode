package main

import (
	"fmt"
	"sort"
)

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
	sort.Ints(nums)
	numsAll := [][]int{}
	n := len(nums)
	if len(nums) ==0 {
		return numsAll
	}
	for i:=0; i < len(nums) ; i ++ {
		for j:=0; j < len(nums) ; j ++ {
			//if nums[i] == nums[j] || i == j {
			//	continue
			//}
			tmp := make([]int,n)
			copy(tmp,nums)
			tmp[i],tmp[j] = tmp[j],tmp[i]
			numsAll = append(numsAll,tmp)
		}
	}
	return numsAll
}

func permute(nums []int) [][]int {
		numsAll := make([][]int,0)
		n := len(nums)
		var backtrack func(pos int )

		if n ==  0 {
			return numsAll
		}
		// 对第pos位置填入值
		backtrack = func(pos int ) {
			// 如果填的位置是最后一位, 则已经填完
			if pos == n-1 {
				// 新建slice是因为, nums会改变, 否则添加的都是最后一次修改的结果
				tmp := make([]int, n)
				copy(tmp,nums)
				numsAll = append(numsAll,tmp)
			}
			// [0:pos] 位置已经确定, 继续填入[pos,n-1]
			for i := pos; i < n ; i ++ {
				nums[i], nums[pos] = nums[pos], nums[i]
				backtrack(pos + 1)
				nums[i], nums[pos] = nums[pos], nums[i]
			}
		}
		backtrack(0)
		return numsAll
}

func main() {
	//fmt.Println(permuteUnique([]int{1,2,3}))
	fmt.Println(permute([]int{1,3,4}))
}