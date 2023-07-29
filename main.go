package main

import "fmt"

func main() {
	C := func(i value) Entry { return constantEntry(i) }
	Op := func(o opcode) Entry { return opEntry(o) }
	// Make our stack
	s := newStack()
	// Create code to push 1 and 2 on the stack, add them, print the
	// result, and return:
	c := newChunk(C(1), C(2), Op(op_add), Op(op_ret))
	// Run the code in a new VM:
	vm := newVM(c)
	// Run the VM
	fmt.Println(vm.run(s))
}
