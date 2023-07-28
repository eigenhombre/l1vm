package main

import "fmt"

func main() {
	// Make our stack
	s := newStack()
	// Create a chunk of code
	c := newChunk()
	// Create code to push 1 and 2 on the stack, add them, print the
	// result, and return:
	c.add(constantEntry(1))
	c.add(constantEntry(2))
	c.add(opEntry(op_add))
	c.add(opEntry(op_print))
	c.add(opEntry(op_ret))
	// Create a VM
	vm := newVM(c)
	// Run the VM
	vm.run(s)
	fmt.Println("OK")
}
