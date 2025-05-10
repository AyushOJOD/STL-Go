package containers

import "fmt"

// Pair holds two values of potentially different types.
type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

// NewPair creates a new pair with the provided values.
func NewPair[T1 any, T2 any](first T1, second T2) Pair[T1, T2] {
	return Pair[T1, T2]{First: first, Second: second}
}

// Swap returns a new pair with First and Second swapped.
func (p Pair[T1, T2]) Swap() Pair[T2, T1] {
	return NewPair(p.Second, p.First)
}

// String returns the string representation of the pair.
func (p Pair[T1, T2]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

