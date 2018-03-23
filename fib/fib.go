package fib

// NewFibState initializes to iterate from the beginning of the sequence
func NewFibState() FibState {
	return FibState{A: 0, B: 0}
}

// FibState is an iterator used to iterate the fibonacci sequence
type FibState struct {
	B uint64
	A uint64
}

func (s *FibState) Next() uint64 {
	// calculate the next fibonacci number
	c := s.A + s.B

	// if both of our current numbers are 0, the next is actually a 1
	if c == 0 {
		c = 1
	}

	// update the iterator state
	s.A = s.B
	s.B = c

	// return the new number
	return c
}
