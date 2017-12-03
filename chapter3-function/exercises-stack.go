package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func main() {
	// s := stack{3, [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	s := stack{}
	s.push(10)
	s.push(20)
	s.push(30)
	s.push(40)
	s.push(50)
	s.push(60)
	s.push(70)
	s.push(80)
	s.pop()
	s.pop()
	s.push(100)
	fmt.Println(s)
}

func (s *stack) push(item int) {
	s.data[s.i] = item
	s.i++
	// fmt.Println(s)
}

func (s *stack) pop() {
	s.i--
	item := s.data[s.i]
	s.data[s.i] = 0

	fmt.Printf("Remove %d from stack\n", item)
}

func (s stack) String() (str string) {
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return
}
