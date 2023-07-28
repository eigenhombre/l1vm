package main

// A chunk consists of opcodes or constants, represented
// as "value" types:
const (
	typ_opcode = iota
	typ_constant
)

type Entry struct {
	typ int
	opcode
	value
}

type chunk struct {
	entries []Entry
	index   int
}

func newChunk() *chunk {
	return &chunk{[]Entry{}, 0}
}

func (c *chunk) add(e Entry) {
	c.entries = append(c.entries, e)
}

func opEntry(o opcode) Entry {
	return Entry{typ_opcode, o, 0}
}

func constantEntry(v value) Entry {
	return Entry{typ_constant, 0, v}
}

func (c *chunk) next() Entry {
	e := c.entries[c.index]
	c.index++
	return e
}

func (c *chunk) rewind() {
	c.index = 0
}

func (c *chunk) hasMore() bool {
	return c.index < len(c.entries)
}
