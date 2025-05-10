package algorithms

import (
	"slices"
	"sort"

	"golang.org/x/exp/constraints"
)

// Reverse reverses the elements in the given slice in-place.
// @param data the slice to reverse
// @param from the start index (inclusive)
// @param to the end index (exclusive)
func Reverse[T any](data []T, from, to int) {
	if from < 0 || to > len(data) || from >= to {
		return
	}

	for i, j := from, to-1; i < j; i, j = i+1, j-1 {
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


// Sort sorts data[from:to] in-place using default comparison.
// @param data the slice to partially sort
// @param from the starting index (inclusive)
// @param to the ending index (exclusive)
func Sort[T constraints.Ordered](data []T, from, to int) {
	if from < 0 || to > len(data) || from >= to {
		return
	}
	slices.Sort(data[from:to])
}
// SortComp sorts the elements in data[from:to] using the provided comparison function.
// @param data the slice to partially sort
// @param from the starting index (inclusive)
// @param to the ending index (exclusive)
// @param less the comparison function
func SortComp[T any](data []T, from, to int, less func(a, b T) bool) {
	// Validate bounds
	if from < 0 || to > len(data) || from >= to {
		return
	}

	sort.Slice(data[from:to], func(i, j int) bool {
		return less(data[from+i], data[from+j])
	})
}

// IsSorted checks if data[from:to] is sorted in ascending order.
// @param data the slice to check partially
// @param from the starting index (inclusive)
// @param to the ending index (exclusive)
// @return true if sorted, false otherwise
func IsSorted[T constraints.Ordered](data []T, from, to int) bool {
	// Validate bounds
	if from < 0 || to > len(data) || from >= to {
		return true // An empty or one-element range is trivially sorted
	}

	return slices.IsSorted(data[from:to])
}

// GetMax returns the maximum value and its index (absolute) in data[from:to].
// @param data the slice to search partially
// @param from the starting index (inclusive)
// @param to the ending index (exclusive)
// @return maximum value and its absolute index in data
func GetMax[T constraints.Ordered](data []T, from, to int) (T, int) {
	if from < 0 || to > len(data) || from >= to {
		panic("invalid range")
	}

	maxValue := data[from]
	maxIndex := from

	for i := from + 1; i < to; i++ {
		if data[i] > maxValue {
			maxValue = data[i]
			maxIndex = i
		}
	}
	return maxValue, maxIndex
}

// GetMin returns the minimum value and its index (absolute) in data[from:to].
// @param data the slice to search partially
// @param from the starting index (inclusive)
// @param to the ending index (exclusive)
// @return minimum value and its absolute index in data
func GetMin[T constraints.Ordered](data []T, from, to int) (T, int) {
	if from < 0 || to > len(data) || from >= to {
		panic("invalid range")
	}

	min := data[from]
	minIndex := from

	for i := from + 1; i < to; i++ {
		if data[i] < min {
			min = data[i]
			minIndex = i
		}
	}
	return min, minIndex
}


// Accumulate returns the sum of data[from:to].
// @param data the slice to accumulate partially
// @param from the starting index (inclusive)
// @param to the ending index (exclusive)
// @return the sum of the subrange
func Accumulate[T constraints.Ordered](data []T, from, to int) T {
	if from < 0 || to > len(data) || from >= to {
		var zero T
		return zero
	}

	sum := data[from]
	for i := from + 1; i < to; i++ {
		sum += data[i]
	}
	return sum
}
