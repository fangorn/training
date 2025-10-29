package deadfish

// https://www.codewars.com/kata/51e0007c1f9378fa810002a9
type program struct {
	state  int
	output []int
}

func newProgram() *program {
	return &program{
		state:  0,
		output: make([]int, 0),
	}
}

func (p *program) inc() {
	p.state++
}

func (p *program) dec() {
	p.state--
}

func (p *program) sqr() {
	p.state *= p.state
}

func (p *program) out() {
	p.output = append(p.output, p.state)
}

func Parse(data string) []int {
	p := newProgram()
	for _, op := range data {
		switch op {
		case 'i':
			p.inc()
		case 'd':
			p.dec()
		case 's':
			p.sqr()
		case 'o':
			p.out()
		default:
			continue
		}
	}

	return p.output
}
