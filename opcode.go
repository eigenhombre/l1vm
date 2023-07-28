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
