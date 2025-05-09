package containers

import (
	"STL-go/utils"
	"fmt"
)

// Vector is a generic dynamic array container similar to C++'s std::vector.
type Vector[T any] struct {
	data []T
}

// NewVector creates and returns a new empty vector.
// @return pointer to a new Vector[T]
func NewVector[T any]() *Vector[T] {
	return &Vector[T]{
		data: make([]T, 0),
	}
}

// Assign fills the vector with 'size' elements, each set to 'value'.
// @param size number of elements to assign
// @param value value to fill each element with
func (v *Vector[T]) Assign(size int, value T) {
	v.data = make([]T, size)
	for i := range v.data {
		v.data[i] = value
	}
}

// Capacity returns the current capacity of the vector (not size).
// @return integer representing total allocated space
func (v *Vector[T]) Capacity() int {
	return cap(v.data)
}

// Resize changes the size of the vector to the given size.
// If the size increases, new elements are zero-initialized.
// @param size new size of the vector
func (v *Vector[T]) Resize(size int) {
	if size > cap(v.data) {
		newData := make([]T, size)
		copy(newData, v.data)
		v.data = newData
	} else {
		v.data = v.data[:size]
	}
}

// PushBack adds a new element at the end of the vector.
// @param value value to append
func (v *Vector[T]) PushBack(value T) {
	v.data = append(v.data, value)
}

// PopBack removes and returns the last element of the vector.
// @return the last element and true if successful, false if empty
func (v *Vector[T]) PopBack() (T, bool) {
	if len(v.data) == 0 {
		var zero T
		return zero, false
	}
	lastIndex := len(v.data) - 1
	lastValue := v.data[lastIndex]
	v.data = v.data[:lastIndex]
	return lastValue, true
}

// Size returns the number of elements in the vector.
// @return current number of elements
func (v *Vector[T]) Size() int {
	return len(v.data)
}

// At returns the element at the given index.
// @param index position to access
// @return element and true if valid, false otherwise
func (v *Vector[T]) At(index int) (T, bool) {
	if index < 0 || index >= len(v.data) {
		var zero T
		return zero, false
	}
	return v.data[index], true
}

// Clear removes all elements from the vector.
func (v *Vector[T]) Clear() {
	v.data = v.data[:0]
}

// Begin returns a pointer to the first element.
// @return pointer to the first element
func (v *Vector[T]) Begin() *T {
	return &v.data[0]
}

// End returns a pointer to the position after the last element (unsafe).
// @return pointer to one-past-the-end (invalid dereference)
func (v *Vector[T]) End() *T {
	return &v.data[len(v.data)]
}

// IsEmpty returns true if the vector contains no elements.
// @return true if empty, false otherwise
func (v *Vector[T]) IsEmpty() bool {
	return len(v.data) == 0
}

// Print prints the contents of the vector.
func (v *Vector[T]) Print() {
	fmt.Println(v.data)
}

// Insert inserts a value at the given index, shifting elements right.
// @param index position to insert at
// @param value value to insert
func (v *Vector[T]) Insert(index int, value T) {
	if index < 0 || index > len(v.data) {
		return
	}

	v.data = append(v.data, value)

	if index == len(v.data)-1 {
		return
	}

	copy(v.data[index+1:], v.data[index:len(v.data)-1])
	v.data[index] = value
}

// RemoveAt removes the element at the specified index.
// @param index position to remove from
func (v *Vector[T]) RemoveAt(index int) {
	if !utils.IsValidIndex[T](index, len(v.data)) {
		return
	}
	v.data = append(v.data[:index], v.data[index+1:]...)
}

// Front returns the first element of the vector.
// @return the first element and true if vector is not empty, false otherwise
func (v *Vector[T]) Front() (T, bool) {
	if v.IsEmpty() {
		var zero T
		return zero, false
	}
	return v.data[0], true
}

// Back returns the last element of the vector.
// @return the last element and true if vector is not empty, false otherwise
func (v *Vector[T]) Back() (T, bool) {
	if v.IsEmpty() {
		var zero T
		return zero, false
	}
	return v.data[len(v.data)-1], true
}

// Data returns the underlying slice of the vector.
// @return the underlying slice
func (v *Vector[T]) Data() []T {
	return v.data
}