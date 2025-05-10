package containers_examples

import (
	"fmt"

	"github.com/AyushOJOD/stl-go/containers"
)

func ExampleQueue() {
	q := containers.NewQueue[string]()

	q.Push("A")
	q.Push("B")
	q.Push("C")

	front, _ := q.Front()
	fmt.Println("Front:", front)

	fmt.Println("Queue contents:", q.Data())

	for !q.IsEmpty() {
		val, _ := q.Pop()
		fmt.Println("Popped:", val)
	}

	fmt.Println("IsEmpty?", q.IsEmpty())

	// Output:
	// Front: A
	// Queue contents: [A B C]
	// Popped: A
	// Popped: B
	// Popped: C
	// IsEmpty? true
}