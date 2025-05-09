package tests

import (
	"STL-go/containers"
	"testing"
)

func TestQueueBasic(t *testing.T) {
	q := containers.NewQueue[int]()

	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}

	q.Push(10)
	q.Push(20)
	q.Push(30)

	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}

	front, ok := q.Front()
	if !ok || front != 10 {
		t.Errorf("Expected front 10, got %v", front)
	}

	val, ok := q.Pop()
	if !ok || val != 10 {
		t.Errorf("Expected popped value 10, got %v", val)
	}

	_, _ = q.Pop()
	_, _ = q.Pop()
	_, ok = q.Pop()
	if ok {
		t.Errorf("Expected false on pop from empty queue")
	}
}
