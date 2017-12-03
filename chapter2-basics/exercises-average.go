package main

import (
	"fmt"
)

func main() {
	var xs = []float64{100, 200, 300, 400.000}
	fmt.Printf("Value %v\n", xs)
	fmt.Println("len()", len(xs))

	var total float64
	for _, v := range xs {
		total += v
	}

	avgVal := total / float64(len(xs))
	fmt.Println("Total", total)
	fmt.Println("Average", avgVal)

}
