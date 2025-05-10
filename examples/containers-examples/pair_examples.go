package containers_examples

import (
	"fmt"

	"github.com/AyushOJOD/stl-go/containers"
)

func ExamplePairs() {
	p := containers.NewPair("age", 25)
	fmt.Println("Pair:", p)

	swapped := p.Swap()
	fmt.Println("Swapped:", swapped)
}
