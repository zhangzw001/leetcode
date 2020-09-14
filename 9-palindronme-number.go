package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	palindronme := []byte(strconv.Itoa(x))
	fmt.Println(palindronme)
	start := 0
	end := len(palindronme)-1
	for start < end {
		if palindronme[start] != palindronme[end] {
			fmt.Println(palindronme[start],palindronme[end])
			return false
		}else {
			start++
			end--
		}
	}
	return true
}


func main() {
	fmt.Println(isPalindrome(13231))

}