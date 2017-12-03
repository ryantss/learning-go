package main

import (
	"fmt"
)

func main() {
	fmt.Println("Write a function that takes a variable number of ints and print each integer on a separate line.")

	printVarargs(1, 2, 3, 4, 5, 12, 3, 123123, 1)
}

func printVarargs(vals ...int) {

	for _, v := range vals {
		fmt.Println(v)
	}

}
