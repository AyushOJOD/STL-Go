package tests

import (
	"STL-go/containers"
	"testing"
)

func TestDequeueBasic(t *testing.T) {
	d := containers.NewDequeue[int]()

	// PushFront and PushBack
	d.PushBack(1)
	d.PushFront(0)
	d.PushBack(2)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	// Front and Back
	front, ok := d.Front()
	if !ok || front != 0 {
		t.Errorf("Expected front 0, got %v", front)
	}

	back, ok := d.Back()
	if !ok || back != 2 {
		t.Errorf("Expected back 2, got %v", back)
	}

	// PopFront
	val, ok := d.PopFront()
	if !ok || val != 0 {
		t.Errorf("Expected PopFront 0, got %v", val)
	}

	// PopBack
	val, ok = d.PopBack()
	if !ok || val != 2 {
		t.Errorf("Expected PopBack 2, got %v", val)
	}

	// Remaining element
	val, ok = d.PopFront()
	if !ok || val != 1 {
		t.Errorf("Expected remaining element 1, got %v", val)
	}

	if !d.IsEmpty() {
		t.Errorf("Expected dequeue to be empty")
	}

	// Clear
	d.PushBack(10)
	d.PushBack(20)
	d.Clear()
	if !d.IsEmpty() {
		t.Errorf("Expected dequeue to be empty after Clear()")
	}
}
