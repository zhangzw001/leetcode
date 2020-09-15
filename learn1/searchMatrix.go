package main

import "fmt"

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions/xmlwi1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
//从右上角
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) ==0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	p1,p2 := 0, len(matrix[0])-1
	for p1 <=len(matrix)-1 && p2 >=0 {
		if matrix[p1][p2] == target {
			return true
		}

		if matrix[p1][p2] < target {
			p1++
		}else {
			p2--
		}
	}
	return false
}

//从左下角
func searchMatrix(matrix [][]int, target int) bool {
	// 解决[]列表
	if len(matrix) ==0 {
		return false
	}
	// 解决[[]]列表
	if len(matrix[0]) == 0 {
		return false
	}
	p1,p2:=len(matrix)-1,0
	for p1 >=0 && p2 < len(matrix[p1]) {
		if matrix[p1][p2] == target {
			return true
		}
		if matrix[p1][p2] > target {
			p1--
		}else {
			p2 ++
		}
	}
	return false
}

func main() {
	a := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	//a := [][]int{{-1,3}}
	fmt.Println(a)
	fmt.Println(searchMatrix2(a,3))
}