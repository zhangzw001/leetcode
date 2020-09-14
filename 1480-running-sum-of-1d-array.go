package main

import "fmt"

//https://leetcode-cn.com/problems/running-sum-of-1d-array/

func runningSum(nums []int) []int {
	for i,_ := range nums {
		if i != 0 {
			nums[i] +=nums[i-1]
		}
	}
	return nums
}

func main() {
	n := []int{1,2,3,4}
	fmt.Println(runningSum(n))
}