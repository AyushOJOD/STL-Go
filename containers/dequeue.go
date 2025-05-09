package containers

// Dequeue is a double-ended queue.
// It is a collection of elements that supports adding and removing elements from both ends.
// It is similar to a stack and a queue, but it is more flexible.
type Dequeue[T any] struct {
	data []T
}

func NewDequeue[T any]() *Dequeue[T] {
	return &Dequeue[T]{
		data: []T{},
	}
}

// PushFront adds an element to the front of the dequeue.
// @param value the value to add
// @return pointer to the dequeue
func (d *Dequeue[T]) PushFront(value T) {
	d.data = append([]T{value}, d.data...)
}

// PushBack adds an element to the back of the dequeue.
// @param value the value to add
// @return pointer to the dequeue
func (d *Dequeue[T]) PushBack(value T) {
	d.data = append(d.data, value)
}

// PopFront removes and returns the front element of the dequeue.
// @return the front element and a boolean indicating if the element was removed
func (d *Dequeue[T]) PopFront() (T, bool) {
	if len(d.data) == 0 {
		var zero T
		return zero, false
	}
	front := d.data[0]
	d.data = d.data[1:]
	return front, true
}

// PopBack removes and returns the back element of the dequeue.
// @return the back element and a boolean indicating if the element was removed
func (d *Dequeue[T]) PopBack() (T, bool) {
	if len(d.data) == 0 {
		var zero T
		return zero, false
	}

	back := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return back, true
}

// Front returns the front element of the dequeue.
// @return the front element and a boolean indicating if the element was removed
func (d *Dequeue[T]) Front() (T, bool) {
	if len(d.data) == 0 {
		var zero T
		return zero, false
	}
	return d.data[0], true
}

// Back returns the back element of the dequeue.
// @return the back element and a boolean indicating if the element was removed
func (d *Dequeue[T]) Back() (T, bool) {
	if len(d.data) == 0 {
		var zero T
		return zero, false
	}
	return d.data[len(d.data)-1], true
}

// Size returns the number of elements in the dequeue.
// @return the number of elements in the dequeue
func (d *Dequeue[T]) Size() int {
	return len(d.data)
}

// IsEmpty returns true if the dequeue is empty.
// @return true if the dequeue is empty, false otherwise
func (d *Dequeue[T]) IsEmpty() bool {
	return len(d.data) == 0
}

// Clear removes all elements from the dequeue.
func (d *Dequeue[T]) Clear() {
	d.data = []T{}
}

// Data returns the underlying data slice.
// @return the underlying data slice
func (d *Dequeue[T]) Data() []T {
	return d.data
}
