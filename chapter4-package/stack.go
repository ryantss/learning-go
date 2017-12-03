package stack

/*
% mkdir $GOPATH/src/stack
% cp even.go $GOPATH/src/stack
% cd $GOPATH/src/stack
% go build
% go install
*/

import (
	"fmt"
	"strconv"
)

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(item int) {
	s.data[s.i] = item
	s.i++
	// fmt.Println(s)
}

func (s *Stack) Pop() {
	s.i--
	item := s.data[s.i]
	s.data[s.i] = 0

	fmt.Printf("Remove %d from stack\n", item)
}

func (s Stack) String() (str string) {
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return
}
