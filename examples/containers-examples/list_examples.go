package containers_examples

import (
	"fmt"

	"github.com/AyushOJOD/stl-go/containers"
)

func ExampleLists() {
	// Create a new list
	list := containers.NewList[int]()

	// Add elements to the list
	list.PushBack(10)
	list.PushBack(20)
	list.PushFront(5)

	// Display the elements of the list
	fmt.Println("List elements:", list.Data()) // Output: [5 10 20]

	// Pop an element from the front
	frontValue, ok := list.PopFront()
	if ok {
		fmt.Println("Popped from front:", frontValue) // Output: 5
	}

	// Pop an element from the back
	backValue, ok := list.PopBack()
	if ok {
		fmt.Println("Popped from back:", backValue) // Output: 20
	}

	// Display the current elements in the list
	fmt.Println("List elements:", list.Data()) // Output: [10]

	// Front and Back values
	front, ok := list.Front()
	if ok {
		fmt.Println("Front element:", front) // Output: 10
	}

	back, ok := list.Back()
	if ok {
		fmt.Println("Back element:", back) // Output: 10
	}

	// Check size and if empty
	fmt.Println("Size:", list.Size())     // Output: 1
	fmt.Println("Is list empty?", list.IsEmpty()) // Output: false

	// Clear the list
	list.Clear()
	fmt.Println("Is list empty after clear?", list.IsEmpty()) // Output: true
}