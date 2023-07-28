package main

import (
	"testing"
)

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
