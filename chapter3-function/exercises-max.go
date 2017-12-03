package main

import "fmt"

func main() {
	fmt.Println(max([]int{1, 2, 3, 4, 5, -100, 1000, 60}))
}

func max(xs []int) (max int) {
	for _, v := range xs {
		if max < v {
			max = v
		}
	}
	return
}
