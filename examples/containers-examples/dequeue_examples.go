package containers_examples

import (
	"STL-go/containers"
	"fmt"
)

func ExampleDequeue() {
	d := containers.NewDequeue[string]()

	d.PushBack("apple")
	d.PushFront("banana")
	d.PushBack("cherry")

	fmt.Println("Deque Size:", d.Size())

	front, _ := d.Front()
	back, _ := d.Back()

	fmt.Println("Front:", front)
	fmt.Println("Back:", back)

	val, _ := d.PopFront()
	fmt.Println("Popped Front:", val)

	val, _ = d.PopBack()
	fmt.Println("Popped Back:", val)

	fmt.Println("Remaining:", d.Data())

	// Output:
	// Deque Size: 3
	// Front: banana
	// Back: cherry
	// Popped Front: banana
	// Popped Back: cherry
	// Remaining: [apple]
}