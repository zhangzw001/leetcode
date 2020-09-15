package main


func fastSort(l []int ) {
	start,end := 0, len(l)-1
	p := l[start]
	for start < end {

		if p > l[end] {
			l[end],l[start] = l[start],l[end]
		}else {
			end--
		}
	}
}
