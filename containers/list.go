package containers

// Node is a single element in the list.
// It contains a value and a pointer to the next node.
type Node[T any] struct {
	value T
	next  *Node[T]
}

// List is a collection of elements that are linked together.
// It contains a pointer to the head and tail nodes and the size of the list.
type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// NewList creates a new list.
// @return pointer to the list
func NewList[T any]() *List[T] {
	return &List[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// PushBack adds a new element to the end of the list.
// @param value the value to add
// @return pointer to the list
func (l *List[T]) PushBack(value T) {
	newNode := &Node[T]{value: value}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

// PushFront adds a new element to the front of the list.
// @param value the value to add
// @return pointer to the list
func (l *List[T]) PushFront(value T) {
	newNode := &Node[T]{value: value, next: l.head}
	l.head = newNode
	if l.tail == nil {
		l.tail = newNode
	}
	l.size++
}

// PopFront removes and returns the front element of the list.
// @return the front element and true if successful, zero value and false if empty
// @return pointer to the list
func (l *List[T]) PopFront() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	value := l.head.value
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return value, true
}

// PopBack removes and returns the last element of the list.
// @return the last element and true if successful, zero value and false if empty
// @return pointer to the list
func (l *List[T]) PopBack() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}

	if l.head == l.tail {
		value := l.head.value
		l.head = nil
		l.tail = nil
		l.size--
		return value, true
	}

	current := l.head
	for current.next != l.tail {
		current = current.next
	}
	value := l.tail.value
	l.tail = current
	l.tail.next = nil
	l.size--
	return value, true
}

// Front returns the front element of the list.
// @return the front element and true if the list is not empty, zero value and false otherwise
func (l *List[T]) Front() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	return l.head.value, true
}

// Back returns the last element of the list.
// @return the last element and true if the list is not empty, zero value and false otherwise
func (l *List[T]) Back() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}
	return l.tail.value, true
}

// Size returns the number of elements in the list.
// @return the number of elements
func (l *List[T]) Size() int {
	return l.size
}

// IsEmpty checks if the list is empty.
// @return true if the list is empty, false otherwise
func (l *List[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear removes all elements from the list.
func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Data returns the underlying elements as a slice.
// @return the list elements as a slice
func (l *List[T]) Data() []T {
	var data []T
	for current := l.head; current != nil; current = current.next {
		data = append(data, current.value)
	}
	return data
}