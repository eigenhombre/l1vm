package main

type vm struct {
	code *chunk
	ip   int
}

func newVM(c *chunk) *vm {
	return &vm{code: c, ip: 0}
}

func (v *vm) execOpcode(o opcode, s *stack) *value {
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
		v, err := s.pop()
		if err != nil {
			panic(err) // fixme
		}
		// End of processing may be signalled by returning a value:
		return &v
	}
	return nil
}

func (v *vm) run(s *stack) value {
	for v.code.hasMore() {
		e := v.code.next()
		switch e.typ {
		case typ_opcode:
			ret := v.execOpcode(e.opcode, s)
			if ret != nil {
				return *ret
			}
		case typ_constant:
			s.push(e.value)
		}
	}
	return 0
}
