package main

import "fmt"

// A stack consists of values

type stack []value

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(v value) {
	*s = append(*s, v)
}

func (s *stack) pop() (value, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("empty stack")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}
