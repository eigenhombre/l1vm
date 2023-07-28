package main

type vm struct {
	code *chunk
	ip   int
}

func newVM(c *chunk) *vm {
	return &vm{code: c, ip: 0}
}

func (v *vm) execOpcode(o opcode, s *stack) {
	switch o {
	case op_add:
		b, err := s.pop()
		if err != nil {
			panic(err) // fixme
		}
		a, err := s.pop()
		if err != nil {
			panic(err) // fixme
		}
		s.push(a + b)
	case op_print:
		v, err := s.pop()
		if err != nil {
			panic(err) // fixme
		}
		println(v)
	case op_ret:
		return
	}
}
func (v *vm) run(s *stack) {
	for v.code.hasMore() {
		e := v.code.next()
		switch e.typ {
		case typ_opcode:
			v.execOpcode(e.opcode, s)
		case typ_constant:
			s.push(e.value)
		}
	}
}
