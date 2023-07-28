package main

const (
	op_ret = iota
	op_add
	op_push
	op_print // this will probably go away or change
)

type opcode byte

func (o opcode) String() string {
	switch o {
	case op_ret:
		return "RET"
	case op_add:
		return "ADD"
	case op_push:
		return "PUSH"
	case op_print:
		return "PRINT"
	default:
		return "UNKNOWN_OPCODE"
	}
}
