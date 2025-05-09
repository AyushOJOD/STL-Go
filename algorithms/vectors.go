package algorithms

import (
	"slices"
	"sort"

	"golang.org/x/exp/constraints"
)

// Reverse reverses the elements in the given slice in-place.
// @param data the slice to reverse
func Reverse[T any](data []T) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// ForEach applies a function to each element in the slice.
// @param data the slice to iterate
// @param fn the function to apply
func ForEach[T any](data []T, fn func(T)) {
	for _, item := range data {
		fn(item)
	}
}


// Sort sorts the slice using Go's default comparison operators.
// This requires T to be ordered (int, float64, string, etc.).
// @param data the slice to sort
func Sort[T constraints.Ordered](data []T) {
	slices.Sort(data)
}

// SortFunc sorts the slice using a custom comparison function.
// @param data the slice to sort
// @param less comparator function: returns true if a < b
func SortFunc[T any](data []T, less func(a, b T) bool) {
	sort.Slice(data, func(i, j int) bool {
		return less(data[i], data[j])
	})
}

// IsSorted checks if the slice is sorted in ascending order.
// @param data the slice to check
// @return true if sorted, false otherwise
func IsSorted[T constraints.Ordered](data []T) bool {
	return slices.IsSorted(data)
}

// GetMax returns the maximum value and its index in the slice.
// @param data the slice to search
// @return maximum value and its index
func GetMax[T constraints.Ordered] (data []T) (T, int) {
	
	maxValue := data[0]
	maxIndex := 0

	for i, item := range data {
		if item > maxValue {
			maxValue = item	
			maxIndex = i
		}
	}

	return maxValue, maxIndex
}

// GetMin returns the minimum value and its index in the slice.
// @param data the slice to search
// @return minimum value and its index
func GetMin[T constraints.Ordered] (data []T) (T, int) {
	min := data[0]
	minIndex := 0

	for i, item := range data {	
		if item < min {
			min = item
			minIndex = i
		}
	}
	return min, minIndex
}

// Accumulate returns the sum of all elements in the slice.
// @param data the slice to accumulate
// @return the sum of all elements
func Accumulate[T constraints.Ordered] (data []T) T {

	if len(data) == 0 {
		var zero T
		return zero
	}

	sum := data[0]
	for i := 1; i < len(data); i++ {
		sum += data[i]
	}
	return sum
}

