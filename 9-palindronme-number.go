package main

import (
	"fmt"
	"strconv"
)

func isPalindrome2(x int) bool {
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

func isPalindrome(x int) bool {
	palindronme := strconv.Itoa(x)
	l := len(palindronme)-1
	end := l/2
	fmt.Println(end)
	for i:=0; i < end ; i++ {
		if palindronme[i] != palindronme[l-i] {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(isPalindrome(13231))

}