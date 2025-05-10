package tests

import (
	"testing"

	"github.com/AyushOJOD/stl-go/containers"
)

func TestList(t *testing.T) {
	// Create a new list
	list := containers.NewList[int]()

	// Test PushBack
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}

	// Test PushFront
	list.PushFront(5)
	if list.Size() != 4 {
		t.Errorf("Expected size 4, got %d", list.Size())
	}

	// Test PopFront
	frontValue, ok := list.PopFront()
	if !ok || frontValue != 5 {
		t.Errorf("Expected 5 from PopFront, got %v", frontValue)
	}

	// Test PopBack
	backValue, ok := list.PopBack()
	if !ok || backValue != 30 {
		t.Errorf("Expected 30 from PopBack, got %v", backValue)
	}

	// Test Front and Back
	front, ok := list.Front()
	if !ok || front != 10 {
		t.Errorf("Expected 10 from Front, got %v", front)
	}

	back, ok := list.Back()
	if !ok || back != 20 {
		t.Errorf("Expected 20 from Back, got %v", back)
	}

	// Test Size and IsEmpty
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}

	if list.IsEmpty() {
		t.Errorf("Expected list to be not empty")
	}

	// Test Clear
	list.Clear()
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty after clear")
	}
}

func TestEmptyListOperations(t *testing.T) {
	// Create an empty list
	list := containers.NewList[int]()

	// Test PopFront and PopBack on an empty list
	if _, ok := list.PopFront(); ok {
		t.Error("Expected PopFront to fail on empty list")
	}
	if _, ok := list.PopBack(); ok {
		t.Error("Expected PopBack to fail on empty list")
	}

	// Test Front and Back on an empty list
	if _, ok := list.Front(); ok {
		t.Error("Expected Front to fail on empty list")
	}
	if _, ok := list.Back(); ok {
		t.Error("Expected Back to fail on empty list")
	}

	// Test IsEmpty on an empty list
	if !list.IsEmpty() {
		t.Error("Expected IsEmpty to return true for empty list")
	}
}