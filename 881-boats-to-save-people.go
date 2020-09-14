package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left , right := 0, len(people)-1
	num :=0
	for left <= right {
		if people[left] + people[right] <= limit {
			left++
		}
		right --
		num ++
	}
	return num
}
func numRescueBoats2(people []int, limit int) int {
	var bucket = make([]int, limit+1)
	for _, p := range people {
		bucket[p]++
	}
	var ans int
	for i:=len(bucket)-1; i>0; i-- {
		for j:=limit-i; bucket[i]>0 && j>0; j-- {
			if bucket[j] == 0 {
				continue
			}
			if i == j {
				ans += bucket[i]/2
				bucket[i] = bucket[i]%2
				continue
			}
			if bucket[j] <= bucket[i] {
				ans += bucket[j]
				bucket[i] -= bucket[j]
				bucket[j] = 0
				continue
			}
			ans += bucket[i]
			bucket[j] -= bucket[i]
			bucket[i] = 0
		}
		if bucket[i] > 0 {
			ans += bucket[i]
			bucket[i] = 0
		}
	}
	return ans
}

func main() {
	fmt.Println(numRescueBoats2([]int{3,2,2,1},3))
}