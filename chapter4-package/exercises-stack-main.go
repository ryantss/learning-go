package main

import (
	"fmt"
	"stack"
)

func main() {
	// s := stack{3, [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	s := stack.Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Push(50)
	s.Push(60)
	s.Push(70)
	s.Push(80)
	s.Pop()
	s.Pop()
	s.Push(100)
	fmt.Println(s)
}
