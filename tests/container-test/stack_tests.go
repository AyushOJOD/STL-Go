package tests

import (
	"testing"

	"github.com/AyushOJOD/stl-go/containers"
)

func TestStackBasic(t *testing.T) {
	s := containers.NewStack[int]()

	if !s.IsEmpty() {
		t.Error("Expected stack to be empty initially")
	}

	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	top, ok := s.Top()
	if !ok || top != 30 {
		t.Errorf("Expected top 30, got %v", top)
	}

	val, ok := s.Pop()
	if !ok || val != 30 {
		t.Errorf("Expected pop 30, got %v", val)
	}

	val, ok = s.Pop()
	if !ok || val != 20 {
		t.Errorf("Expected pop 20, got %v", val)
	}

	s.Clear()
	if !s.IsEmpty() {
		t.Error("Expected stack to be empty after clear")
	}

	_, ok = s.Pop()
	if ok {
		t.Error("Expected pop on empty stack to return false")
	}
}
