package main

import "testing"

func TestChunk(t *testing.T) {
	c := newChunk()
	if c.hasMore() {
		t.Error("new chunk should not have more")
	}
	c.add(opEntry(op_ret))
	if !c.hasMore() {
		t.Error("chunk should have more")
	}
	c.add(constantEntry(1))
	if !c.hasMore() {
		t.Error("chunk should have more")
	}
	entry := c.next()
	if entry.typ != typ_opcode {
		t.Error("entry should be an opcode")
	}
	if entry.opcode != op_ret {
		t.Error("entry should be an op_ret")
	}
	entry = c.next()
	if entry.typ != typ_constant {
		t.Error("entry should be a constant")
	}
	if c.hasMore() {
		t.Error("chunk should not have more")
	}
}
