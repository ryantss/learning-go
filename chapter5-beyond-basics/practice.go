package main

import (
	"fmt"
)

type NameAge struct {
	name string
	age  int
}

func main() {
	var p *int
	fmt.Printf("%v\n", p)

	var i int
	p = &i
	fmt.Printf("%v\n", p)

	// NameAge
	a := new(NameAge)
	a.name = "Ryan"
	a.age = 42
	fmt.Printf("%v\n", a)
}
