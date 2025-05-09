package containers

// Queue is a generic FIFO (first-in, first-out) container.
type Queue[T comparable] struct {
	data []T
}

// NewQueue creates and returns a new empty queue.
// @return pointer to a new Queue[T]
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		data: []T{},
	}
}

// Push adds an element to the back of the queue.
// @param value the value to be added
func (q *Queue[T]) Push(value T) {
	q.data = append(q.data, value)
}

// Pop removes and returns the front element of the queue.
// @return the front element and true if successful, zero value and false if empty
func (q *Queue[T]) Pop() (T, bool) {
	if len(q.data) == 0 {
		var zero T
		return zero, false
	}
	front := q.data[0]
	q.data = q.data[1:]
	return front, true
}

// Front returns the front element without removing it.
// @return the front element and true if successful, zero value and false if empty
func (q *Queue[T]) Front() (T, bool) {
	if len(q.data) == 0 {
		var zero T
		return zero, false
	}
	return q.data[0], true
}

// Size returns the number of elements in the queue.
// @return integer count of elements
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// IsEmpty checks if the queue is empty.
// @return true if the queue is empty, false otherwise
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Clear removes all elements from the queue.
// @return pointer to the queue
func (q *Queue[T]) Clear() {
	q.data = []T{}
}

// Data returns the underlying data slice.
// @return the underlying data slice
func (q *Queue[T]) Data() []T {
	return q.data
}