package main

const (
	op_nop = iota
	op_ret
	op_push
	op_pop
	op_const
	op_jump
	op_plus
	op_times
	op_negate
	op_divide
)

type opcode byte

func (o opcode) String() string {
	switch o {
	case op_nop:
		return "NOP"
	case op_ret:
		return "RET"
	case op_push:
		return "PUSH"
	case op_pop:
		return "POP"
	case op_const:
		return "CONST"
	case op_jump:
		return "JUMP"
	case op_plus:
		return "PLUS"
	case op_times:
		return "TIMES"
	case op_negate:
		return "NEGATE"
	case op_divide:
		return "DIVIDE"
	default:
		return "UNKNOWN_OPCODE"
	}
}

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
