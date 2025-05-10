package containers_examples

import (
	"fmt"

	"github.com/AyushOJOD/stl-go/containers"
)

func ExampleStack() {
	s := containers.NewStack[string]()

	s.Push("A")
	s.Push("B")
	s.Push("C")

	fmt.Println("Size:", s.Size())

	top, _ := s.Top()
	fmt.Println("Top:", top)

	for !s.IsEmpty() {
		val, _ := s.Pop()
		fmt.Println("Popped:", val)
	}

	// Output:
	// Size: 3
	// Top: C
	// Popped: C
	// Popped: B
	// Popped: A
}
