package containers

// Stack is a generic stack container implementing LIFO (Last-In, First-Out).
type Stack[T any] struct {
	data []T
}

// NewStack creates and returns a new empty stack.
// @return pointer to a new Stack[T]
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: []T{},
	}
}

// Push adds an element to the top of the stack.
// @param value the value to be pushed
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top element of the stack.
// @return the top element and true if successful, or zero value and false if empty
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, true
}

// Top returns the top element without removing it from the stack.
// @return the top element and true if successful, or zero value and false if empty
func (s *Stack[T]) Top() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Size returns the number of elements in the stack.
// @return the number of elements
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// IsEmpty checks whether the stack is empty.
// @return true if the stack is empty, false otherwise
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	s.data = []T{}
}

// Data returns the underlying slice of the stack (read-only use recommended).
// @return the slice representing the stack contents
func (s *Stack[T]) Data() []T {
	return s.data
}
