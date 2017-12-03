package main

import "fmt"

func main() {
	forLoop()
	gotoLoop()
	loopWithArray()
}

func forLoop() {
	fmt.Println("Using 'for' loop")
	for i := 0; i < 10; i++ {
		fmt.Println("Counter", i)
	}
}

func gotoLoop() {
	fmt.Println("Using 'goto' for loop")
	i := 0
Loop:
	if i < 10 {
		fmt.Println("Counter", i)
		i++
		goto Loop
	}

}

func loopWithArray() {
	fmt.Println("Loop with array")
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	for _, counter := range arr {
		fmt.Println("Counter", counter)
	}
	fmt.Println("End of arrayLoop")
}
