package main

const (
	op_ret = iota
)

type opcode byte

func (o opcode) String() string {
	switch o {
	case op_ret:
		return "RET"
	default:
		return "UNKNOWN_OPCODE"
	}
}
