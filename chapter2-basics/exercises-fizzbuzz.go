package main

import "fmt"

const (
	FIZZ = 3
	BUZZ = 5
)

func main() {
	fmt.Println("Write a program that prints the numbers from 1 to 100. But for multiples of three print, 'Fizz' instead of the number, and for multiples of five, print 'Buzz'. For numbers which are multiples of both three and five, print 'FizzBuzz'.")
	ifElseMethod()
	// switchMethod()
}

func ifElseMethod() {
	fmt.Println("IF-ELSE METHOD...")
	for i := 1; i <= 100; i++ {
		if i%FIZZ == 0 && i%BUZZ == 0 {
			fmt.Println("FizzBuzz")
		} else if i%FIZZ == 0 {
			fmt.Println("Fizz")
		} else if i%BUZZ == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

/*
ANSWER
func main() {
	const (
		FIZZ = 3 1
		BUZZ = 5
	)
	var p bool                 2
	for i := 1; i < 100; i++ { 3
		p = false
		if i%FIZZ == 0 { 4
			fmt.Printf("Fizz")
			p = true
		}
		if i%BUZZ == 0 { 5
			fmt.Printf("Buzz")
			p = true
		}
		if !p { 6
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}
*/

// func switchMethod() {
// 	fmt.Println("SWITCH METHOD...")
// 	for i := 1; i < 100; i++ {
// 		var s string
// 		switch num := i; {
// 		case num%3 == 0:
// 			s += "Fizz"
// 			fallthrough
// 		case num%5 == 0:
// 			s += "Buzz"
// 		default:
// 			s += strconv.Itoa(num)
// 		}
// 		fmt.Println(s)
// 	}
// }
