package stack

import (
	"stack"
	"testing"
)

func TestPush(t *testing.T) {
	s := stack.Stack{}
	s.Push(123)
	s.Push(456)

	x := s.Pop()
	y := s.Pop()

	if x != 456 {
		t.Log("x (the first pop) should be 456")
		t.Fail()
	}

	if y != 123 {
		t.Log("x (the second pop) should be 123")
		t.Fail()
	}
}
