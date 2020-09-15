package main

import "fmt"

func maxArea1(height []int) int {
	var maxh, maxw, maxv int
	n := len(height)
	for i := range height {
		for j := i ; j < n ; j ++ {
			if height[i] < height[j] {
				maxh = height[i]
			}else {
				maxh = height[j]
			}
			maxw = j-i
			tmp := maxh * maxw
			if maxv < tmp {
				maxv = tmp
			}
		}
	}
	return maxv
}

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}


func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}