package containers_examples

import (
	"STL-go/containers"
	"fmt"
)

func ExampleQueue() {
	q := containers.NewQueue[string]()

	q.Push("apple")
	q.Push("banana")
	q.Push("cherry")

	// Front element
	front, _ := q.Front()
	fmt.Println("Front:", front)

	// Pop elements
	for i := 0; i < 3; i++ {
		val, ok := q.Pop()
		if ok {
			fmt.Println("Popped:", val)
		}
	}

	// Try popping from empty queue
	_, ok := q.Pop()
	fmt.Println("Empty pop?", ok)

	// Output:
	// Front: apple
	// Popped: apple
	// Popped: banana
	// Popped: cherry
	// Empty pop? false
}
