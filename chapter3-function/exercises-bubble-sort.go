package main

import "fmt"

func main() {
	n := []int{5, 3, 2, 1, 10}
	// n := []int{5, -1, 0, 12, 3, 5}
	// bubbleSort(n)
	// fmt.Println(n)

	bubblesortAnswer(n)
	fmt.Println("Answer", n)
}

func bubblesortAnswer(n []int) {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			fmt.Printf("[i:%d j:%d] %v\n", i, j, n)
			fmt.Println(n[j], "<", n[i], "?")
			if n[j] < n[i] {
				// always swap the smallest to the left first and use that to compare against the right.
				n[i], n[j] = n[j], n[i]
			}
		}
		fmt.Println("")
	}
}

func bubbleSort(list []int) {
	swapped := true
Loop:
	if swapped {
		swapped = false
		for i := 1; i < len(list); i++ {
			if list[i-1] > list[i] {
				tmp := list[i-1]
				list[i-1] = list[i]
				list[i] = tmp
				swapped = true
			}
		}
		goto Loop
	}
}
