package tests

import (
	"STL-go/containers"
	"testing"
)

func TestVectorBasic(t *testing.T) {
	v := containers.NewVector[int]()

	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)

	if v.Size() != 3 {
		t.Errorf("Expected size 3, got %d", v.Size())
	}

	val, ok := v.At(1)
	if !ok || val != 2 {
		t.Errorf("Expected value 2 at index 1, got %v", val)
	}

	v.Insert(1, 99)
	val, _ = v.At(1)
	if val != 99 {
		t.Errorf("Expected inserted value 99 at index 1, got %v", val)
	}

	v.RemoveAt(1)
	val, _ = v.At(1)
	if val != 2 {
		t.Errorf("Expected value 2 at index 1 after removal, got %v", val)
	}

	front, ok := v.Front()
	if !ok || front != 1 {
		t.Errorf("Expected front 1, got %v", front)
	}

	back, ok := v.Back()
	if !ok || back != 3 {
		t.Errorf("Expected back 3, got %v", back)
	}

	v.Assign(4, 42)
	if v.Size() != 4 {
		t.Errorf("Expected size 4 after assign, got %d", v.Size())
	}

	v.Clear()
	if !v.IsEmpty() {
		t.Errorf("Expected vector to be empty after clear")
	}
}
