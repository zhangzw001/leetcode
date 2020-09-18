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
	n := len(nums)
	numsAll := make([][]int,0)
	if n ==0 {
		return numsAll
	}

	var backtrack  func(pos int )
	backtrack = func(pos int ) {
		if pos == n-1 {
			tmp := make([]int,n)
			copy(tmp,nums)
			numsAll = append(numsAll,tmp)
			fmt.Println(tmp)
			return
		}

		for i := pos; i < n ; i++ {
			if pos != i && nums[pos] == nums[i] {
				continue
			}
			nums[i] ,nums[pos] = nums[pos],nums[i]
			backtrack(pos+1)
			//nums[i] ,nums[pos] = nums[pos],nums[i]
		}
		for i := n - 1; i > pos; i-- {
			nums[i] ,nums[pos] = nums[pos],nums[i]
		}
	}
	backtrack(0)
	return numsAll
}



func permuteUnique2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	result := [][]int{}
	sort.Ints(nums)
	helper(nums, 0, &result)

	return result
}

// 回溯函数实现
// i表示本次函数需要放置的元素位置
func helper(nums []int, i int, result *[][]int) {
	n := len(nums)
	if i == n-1 {
		tmp := make([]int, n)
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}
	// nums[0:i]是已经决定的部分，nums[i:]是待决定部分，同时待选元素也都在nums[i:]
	for k := i; k < n; k++ {
		// 跳过重复数字
		if k != i && nums[k] == nums[i] {
			continue
		}
		nums[k], nums[i] = nums[i], nums[k]
		helper(nums, i+1, result)
	}
	// 还原状态
	for k := n - 1; k > i; k-- {
		nums[i], nums[k] = nums[k], nums[i]
	}
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
			for i := pos; i < n ; i ++ {
				// 回溯+递归
				nums[i], nums[pos] = nums[pos], nums[i]
				backtrack(pos + 1)
				nums[i], nums[pos] = nums[pos], nums[i]
			}
		}
		backtrack(0)
		return numsAll
}

func main() {
	fmt.Println(permuteUnique([]int{2,2,1,1}))
	//fmt.Println(permuteUnique2([]int{2,2,1,1}))
	//fmt.Println(permuteUnique([]int{1,1,3}))
	//fmt.Println(permuteUnique([]int{1,2}))
	//fmt.Println(permute([]int{1,3,4}))
}