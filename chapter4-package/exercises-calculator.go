package main

import (
	"bufio"
	"fmt"
	"os"
	"stack"
	"strconv"
)

const (
	ADD           = '+'
	SUBTRACT      = '-'
	MULTIPLY      = '*'
	DIVIVE        = '/'
	OPEN_BRACKET  = '('
	CLOSE_BRACKET = ')'
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = new(stack.Stack)

func main() {
	// var formula = "1 + 2"
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}

		for _, c := range s {
			switch {
			case c >= '0' && c <= '9':
				token = token + string(c)
			case c == ' ':
				r, _ := strconv.Atoi(token)
				st.Push(r)
				token = ""
			case c == ADD:
				fmt.Printf("%d\n", st.Pop()+st.Pop())
			case c == MULTIPLY:
				fmt.Printf("%d\n", st.Pop()*st.Pop())
			case c == SUBTRACT:
				p := st.Pop()
				q := st.Pop()
				fmt.Printf("%d\n", q-p)
			case c == 'q':
				return
			default:
				//error
			}
		}
	}
}
