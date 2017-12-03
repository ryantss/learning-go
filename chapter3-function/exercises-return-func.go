package main

import "fmt"

func main() {
	fmt.Println("Write a function that returns a function that performs a +2+2 on integers. Name the function plusTwo. You should then be able do the following:")
	p := plusTwo()
	fmt.Println(p(2))

	x := plusX(20)
	fmt.Println(x(5))
}

func plusTwo() func(int) int {
	return func(x int) int {
		return x + 2
	}
}

func plusX(x int) func(int) int {
	return func(n int) int {
		return x + n
	}
}

// type innerFunction func(int) int

// func plusTwo() innerFunction {
// 	return func(x int) int {
// 		return x + 2
// 	}
// }

// func plusX(x int) innerFunction {
// 	return func(n int) int {
// 		return x + n
// 	}
// }
