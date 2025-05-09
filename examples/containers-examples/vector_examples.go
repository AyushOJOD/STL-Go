package containers_examples

import (
	"STL-go/algorithms"
	"STL-go/containers"
	"fmt"
)

func ExampleVector() {
	// Create a new vector of integers
	vec := containers.NewVector[int]()
	vec.PushBack(10)
	vec.PushBack(20)
	vec.PushBack(30)
	vec.PushBack(40)
	vec.PushBack(50)

	// Example of Reverse
	fmt.Println("Original Vector:", vec.Data())
	algorithms.Reverse(vec.Data())
	fmt.Println("Reversed Vector:", vec.Data())

	// Example of ForEach
	fmt.Println("ForEach Example:")
	algorithms.ForEach(vec.Data(), func(v int) {
		fmt.Println(v)
	})

	// Example of Sort
	fmt.Println("Sorted Vector:")
	algorithms.Sort(vec.Data())
	fmt.Println(vec.Data())

	// Example of IsSorted
	isSorted := algorithms.IsSorted(vec.Data())
	fmt.Printf("Is vector sorted? %v\n", isSorted)

	// Example of SortFunc with a custom comparison (Descending order)
	fmt.Println("Sorted in Descending Order:")
	algorithms.SortFunc(vec.Data(), func(a, b int) bool {
		return a > b
	})
	fmt.Println(vec.Data())

	// Example of GetMax
	maxVal, maxIdx := algorithms.GetMax(vec.Data())
	fmt.Printf("Max Value: %d at Index: %d\n", maxVal, maxIdx)

	// Example of GetMin
	minVal, minIdx := algorithms.GetMin(vec.Data())
	fmt.Printf("Min Value: %d at Index: %d\n", minVal, minIdx)

	// Example of Accumulate (Sum)
	totalSum := algorithms.Accumulate(vec.Data())
	fmt.Printf("Sum of vector elements: %d\n", totalSum)
}
