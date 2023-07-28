package main

import (
	"testing"
)

func TestConstant(t *testing.T) {
	c1 := addConstant(1)
	if getConstant(c1) != 1 {
		t.Error("constant not found")
	}
	c2 := addConstant(2)
	if getConstant(c2) != 2 {
		t.Error("constant not found")
	}
	if getConstant(c1) != 1 {
		t.Error("constant not found")
	}
}
