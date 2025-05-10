package tests

import (
	"testing"

	"github.com/AyushOJOD/stl-go/containers"
)

func TestQueueFull(t *testing.T) {
	q := containers.NewQueue[int]()

	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty initially")
	}

	q.Push(10)
	q.Push(20)
	q.Push(30)

	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}

	if q.IsEmpty() {
		t.Errorf("Expected queue to be non-empty after push")
	}

	front, ok := q.Front()
	if !ok || front != 10 {
		t.Errorf("Expected front 10, got %v", front)
	}

	val, ok := q.Pop()
	if !ok || val != 10 {
		t.Errorf("Expected popped value 10, got %v", val)
	}

	if q.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", q.Size())
	}

	q.Clear()
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty after Clear()")
	}

	data := q.Data()
	if len(data) != 0 {
		t.Errorf("Expected underlying data slice to be empty, got %v", data)
	}
}