package main

import (
	"fmt"
	"testing"
)

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

// simple test to show that the stack is empty
func TestStack(t *testing.T) {
	s := newStack()
	_, err := s.pop()
	if err == nil {
		t.Error("pop from empty stack should return an error")
	}
	s.push(1)
	s.push(2)
	v, err := s.pop()
	if err != nil {
		t.Error("pop should not return an error")
	}
	if v != 2 {
		t.Error("pop should return the value")
	}
	v, err = s.pop()
	if err != nil {
		t.Error("pop should not return an error")
	}
	if v != 1 {
		t.Error("pop should return the value")
	}
}
