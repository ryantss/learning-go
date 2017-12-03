package main

import "fmt"

func main() {
	var p *int
	fmt.Println("%v", p)
	// fmt.Println("%v", *p)

	var i int
	p = &i

	fmt.Println("%v", p)
	fmt.Println("%v", *p)
}
