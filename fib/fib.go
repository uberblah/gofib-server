package fib

import (
	"fmt"
)

// NewFibState initializes to iterate from the beginning of the sequence
func NewFibState() FibState {
	return FibState{A: 0, B: 0}
}

// FibState is an iterator used to iterate the fibonacci sequence
type FibState struct {
	B int
	A int
}

func (s *FibState) Next() int {
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

func ExhaustiveNextFib(n0 int, limit int) (n1 int, err error) {
	// find the next fibonacci number, while ensuring n0 is a fibonacci number
	fibState := NewFibState()
	for {

		// get the next fibonacci number in the sequence
		next := fibState.Next()

		if limit > 0 && next > limit {
			err = fmt.Errorf("Fibonacci number %d exceeded upper limit %d", next, limit)
			return
		}

		// verify that the caller's number still might be a fibonacci number
		if next > n0 {
			err = fmt.Errorf("%d is not a fibonacci number -- it is less than fibonacci number %d and was not found earlier in the sequence", n0, next)
			return
		}

		// if we find the caller's number
		if next == n0 {
			// calculate the *next* fibonacci number
			n1 = fibState.Next()
			return
		}

	}
}
