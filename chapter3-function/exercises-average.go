package main

import (
	"fmt"
)

func main() {
	var xs = []float64{100, 200, 300, 400.000}
	fmt.Printf("Value %v\n", xs)
	// fmt.Println("len()", len(xs))

	avgVal := average(xs)
	fmt.Println("Average", avgVal)

}

func average(numbers []float64) (avg float64) {
	var total float64
	for _, v := range numbers {
		total += v
	}

	avg = total / float64(len(numbers))
	return
}
