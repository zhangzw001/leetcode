package main

import (
	"fmt"
	"sort"
)

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。


说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 

示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions/xmi2l7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums1= append(nums1[:m],nums2[:n]...)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

func merge3(nums1 []int, m int, nums2 []int, n int)  {

	for m > 0 && n > 0 {
		if nums1[m - 1] > nums2[n - 1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
	fmt.Println(nums1)
}

func merge4(nums1 []int, m int, nums2 []int, n int)  {
	n1:= m-1
	n2 :=n-1
	l := m+n-1
	for n2>=0{
		if n1>=0&&nums1[n1]>nums2[n2]{
			nums1[l]= nums1[n1]
			n1--
		}else {
			nums1[l]= nums2[n2]
			n2--
		}
		l--
	}
	fmt.Println(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	//if m == 0 {
	//	nums1=append(nums1[:0],nums2...)
	//}
	left, right , p := m-1, n-1, m+n-1
	for left >= 0 && right >= 0 {
		fmt.Println(nums1[left],nums2[right])
		if nums1[left] > nums2[right] {
			nums1[p] = nums1[left]
			left --
		}else {
			nums1[p] = nums2[right]
			right --
		}
		p --
	}
	fmt.Println(nums1)
}

func main() {
	merge2([]int{2,0},1,[]int{1},1)
	merge3([]int{2,0},1,[]int{1},1)
	merge4([]int{2,0},1,[]int{1},1)
}