package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			fmt.Println(arr, key)
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
			fmt.Println(start,j,end,i )
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

func main() {
	a := []int{3,1,4,2,5,7}
	quickSort(a,0,5)
	fmt.Println(a)
}
