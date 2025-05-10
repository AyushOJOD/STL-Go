package tests

import (
	"testing"

	"github.com/AyushOJOD/stl-go/containers"
)

func TestPairBasic(t *testing.T) {
	p := containers.NewPair[int, string](1, "apple")

	if p.First != 1 || p.Second != "apple" {
		t.Errorf("Expected (1, 'apple'), got (%v, %v)", p.First, p.Second)
	}

	swapped := p.Swap()
	if swapped.First != "apple" || swapped.Second != 1 {
		t.Errorf("Swap failed: got (%v, %v)", swapped.First, swapped.Second)
	}
}
