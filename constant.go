// constant.go
package main

const (
	op_nop = iota
	op_ret
	op_jmp
	op_plus
	op_times
)

// for now, constants are just ints with integer index
// in the constant table
type constant int

var constantTable = []constant{}

func getConstant(i int) constant {
	return constantTable[i]
}

var numConstants = 0

func addConstant(c constant) int {
	constantTable = append(constantTable, c)
	numConstants++
	return numConstants - 1
}