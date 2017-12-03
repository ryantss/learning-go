package main

import "fmt"

func main() {
	f := func(y int) int {
		return y * y
	}

	fmt.Println(Map([]int{1, 2, 3, 4}, f))
}

func Map(list []int, f func(int) int) []int {
	var xs = make([]int, len(list))
	for i, v := range list {
		xs[i] = f(v)
	}
	return xs
}
