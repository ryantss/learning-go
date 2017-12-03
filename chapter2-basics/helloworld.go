package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	// a := [...]int{1, 2, 3, 4, 5}
	// fmt.Println(a[1:5])
	// fmt.Println(a[:])
	// fmt.Println(a[:4])
	// s6 := a[2:4:5]
	// fmt.Println(s6)
	// fmt.Println(len(s6))
	// fmt.Println(cap(s6))
	// fmt.Println(a[1:3:3])

	// a[low : high : max]

	var s = []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(s[1:2:6], len(s[1:2:6]), cap(s[1:2:6]))
	fmt.Println(s[1:2:5], len(s[1:2:5]), cap(s[1:2:5]))
	fmt.Println(s[1:2], len(s[1:2]), cap(s[1:2]))

}
